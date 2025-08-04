package routes

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/handlers"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/repository"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterApprovalRoutes(router fiber.Router, db *gorm.DB) {
	repo := repository.NewApprovalRepository(db)
	service := services.NewApprovalService(repo)
	handler := handlers.NewApprovalHandler(service)

	router.Post("/approve", handler.ApproveExpense)
}
