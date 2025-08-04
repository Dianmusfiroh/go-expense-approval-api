package handlers

import (

    "github.com/gofiber/fiber/v2"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/dto"
	service "github.com/Dianmusfiroh/go-expense-approval-api/internal/services")
type ExpenseHandler struct {
	service service.ExpenseService
}

func NewExpenseHandler(service service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{service}
}

func (h *ExpenseHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateExpenseRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	expense, err := h.service.Create(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(expense)
}

func (h *ExpenseHandler) GetAll(c *fiber.Ctx) error {
	expenses, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(expenses)
}
