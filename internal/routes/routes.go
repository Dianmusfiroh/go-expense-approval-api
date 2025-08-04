package routes

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/handlers"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})
	api := app.Group("/api")
	api.Post("/login", handlers.Login)
	protected := api.Group("/", middleware.JWTProtected())

	RegisterExpenseRoutes(protected, db)
	RegisterUserRoutes(protected, db)
	RegisterApproverRoutes(protected, db)
	RegisterApprovalStageRoutes(protected, db)
	RegisterApprovalRoutes(protected, db)
}
