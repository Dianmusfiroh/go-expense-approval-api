package handlers

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/dto"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ApproverHandler struct {
	Service *services.ApproverService
}

func NewApproverHandler(service *services.ApproverService) *ApproverHandler {
	return &ApproverHandler{Service: service}
}

func (h *ApproverHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateApproverRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	resp, err := h.Service.Create(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *ApproverHandler) GetAll(c *fiber.Ctx) error {
	approvers, err := h.Service.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(approvers)
}

func (h *ApproverHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	approver, err := h.Service.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(approver)
}

func (h *ApproverHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var req dto.UpdateApproverRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	resp, err := h.Service.Update(uint(id), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

func (h *ApproverHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Deleted Data successfully",
		"data":    id,
	})

}
