package routes

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/handlers"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/repository"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterApproverRoutes(router fiber.Router, db *gorm.DB) {
	repo := repository.NewApproverRepository(db)
	svc := services.NewApproverService(repo)
	h := handlers.NewApproverHandler(svc)

	approver := router.Group("/approvers")
	approver.Post("", h.Create)
	approver.Get("", h.GetAll)
	approver.Get("/:id", h.GetByID)
	approver.Delete("/:id", h.Delete)
}
