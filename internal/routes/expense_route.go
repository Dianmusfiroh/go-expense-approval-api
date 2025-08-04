package routes

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/handlers"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/repository"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterExpenseRoutes(router fiber.Router, db *gorm.DB) {
	repo := repository.NewExpenseRepository(db)
	svc := services.NewExpenseService(repo)
	h := handlers.NewExpenseHandler(svc)

	expense := router.Group("/expenses")
	expense.Post("", h.Create)
	expense.Get("", h.GetAll)
}
