package handlers

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/dto"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ApprovalHandler struct {
	Service *services.ApprovalService
}

func NewApprovalHandler(service *services.ApprovalService) *ApprovalHandler {
	return &ApprovalHandler{Service: service}
}

func (h *ApprovalHandler) ApproveExpense(c *fiber.Ctx) error {
	var req dto.ApproveExpenseRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.Service.ApproveExpense(&req); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Pengeluaran berhasil disetujui"})
}
