package routes

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/handlers"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers",
		AllowCredentials: false,
	}))
	api := app.Group("/api")
	api.Post("/login", handlers.Login)
	api.Get("/me", middleware.JWTProtected(), handlers.GetUserLogin) // GET /api/me
	protected := api.Group("/", middleware.JWTProtected())

	RegisterExpenseRoutes(protected, db)
	RegisterUserRoutes(protected, db)
	RegisterApproverRoutes(protected, db)
	RegisterApprovalStageRoutes(protected, db)
	RegisterApprovalRoutes(protected, db)
}
