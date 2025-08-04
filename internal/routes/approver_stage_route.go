package routes

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/handlers"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/repository"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterApprovalStageRoutes(router fiber.Router, db *gorm.DB) {
	repo := repository.NewApprovalStageRepository(db)
	service := services.NewApprovalStageService(repo)
	handler := handlers.NewApprovalStageHandler(service)

	stage := router.Group("/approval-stages")
	stage.Post("", handler.Create)
	stage.Get("", handler.GetAll)
	stage.Get("/:id", handler.GetByID)
	stage.Put("/:id", handler.Update)
	stage.Delete("/:id", handler.Delete)
}
