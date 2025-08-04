package routes

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/handlers"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/repository"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router fiber.Router, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	user := router.Group("/users")

	user.Post("/", userHandler.Create)     // POST /users
	user.Get("/", userHandler.GetAll)      // GET /users
	user.Get("/:id", userHandler.GetByID)  // GET /users/:id
	user.Put("/:id", userHandler.Update)   // PUT /users/:id
	user.Patch("/:id", userHandler.Patch)   // PUT /users/:id
	user.Delete("/:id", userHandler.Delete) // DELETE /users/:id
	user.Delete("/hardDelete/:id", userHandler.HardDelete) // DELETE /users/:id
}
