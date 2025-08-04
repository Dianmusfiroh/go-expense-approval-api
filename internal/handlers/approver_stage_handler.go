package handlers

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/dto"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ApprovalStageHandler struct {
	Service *services.ApprovalStageService
}

func NewApprovalStageHandler(service *services.ApprovalStageService) *ApprovalStageHandler {
	return &ApprovalStageHandler{Service: service}
}

func (h *ApprovalStageHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateApprovalStageRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	stage, err := h.Service.Create(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(stage)
}

func (h *ApprovalStageHandler) GetAll(c *fiber.Ctx) error {
	stages, err := h.Service.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(stages)
}

func (h *ApprovalStageHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	stage, err := h.Service.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(stage)
}

func (h *ApprovalStageHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var req dto.UpdateApprovalStageRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	stage, err := h.Service.Update(uint(id), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(stage)
}

func (h *ApprovalStageHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Deleted data successfully",
	})
}
