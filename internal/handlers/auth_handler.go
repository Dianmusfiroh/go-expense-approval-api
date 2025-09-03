package handlers

import (
	"fmt"
	"time"

	"github.com/Dianmusfiroh/go-expense-approval-api/internal/middleware"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var jwtSecret = []byte("supersecret")

// LoginInput adalah input JSON dari request login
type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	// Parse body input
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Validasi DB dari context
	dbVal := c.Locals("db")
	db, ok := dbVal.(*gorm.DB)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not available",
		})
	}

	// Cari user berdasarkan email
	var user models.User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Cek password (sementara: plain text)
	if input.Password != user.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Buat token JWT
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}


func GetUserLogin(c *fiber.Ctx) error {
	// PERBAIKAN: Logika validasi token sudah dipindah ke middleware.
	// Handler ini sekarang hanya fokus pada tugasnya: mengambil data user.

	// Ambil claims yang sudah divalidasi dari middleware.
	claims, ok := c.Locals("claims").(*middleware.JwtClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not parse claims",
		})
	}
	
	db := c.Locals("db").(*gorm.DB)
	if !ok {
		fmt.Println("[Handler] Gagal mengambil atau mengonversi claims dari context!")
	} else {
		fmt.Printf("[Handler] Claims received: UserID=%d, Role=%s\n", claims.UserID, claims.Role)
	}
	// Cari user berdasarkan id dari claims.
	var user models.User
	if err := db.First(&user, "id = ?", claims.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(fiber.Map{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	})
}



func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Logout successfully",
	})
}

