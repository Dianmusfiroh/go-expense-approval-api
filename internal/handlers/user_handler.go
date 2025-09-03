package handlers

import (
	"strconv"

	"github.com/Dianmusfiroh/go-expense-approval-api/internal/dto"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/services"
	"github.com/gofiber/fiber/v2"

)

type UserHandler struct {
	// Service services.UserService
	Service services.UserServiceInterface
	
}

// func NewUserHandler(service services.UserService) *UserHandler {
// 	return &UserHandler{service}
// }
func NewUserHandler(service services.UserServiceInterface) *UserHandler {
	return &UserHandler{service}
}
func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.Service.Create(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	resp := dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}
func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.Service.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
roleVal, _ := c.Locals("role").(string)
userIDVal, _ := c.Locals("user_id").(uint)


	var responses []dto.UserResponse
	for _, user := range users {
		responses = append(responses, dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			RoleLocal: roleVal,
			UserID:    userIDVal,
		})
	}

	return c.JSON(responses)
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	user, err := h.Service.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	resp := dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	return c.JSON(resp)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedUser, err := h.Service.Update(uint(id), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	resp := dto.UserResponse{
		ID:    updatedUser.ID,
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
		Role:  updatedUser.Role,
	}

	return c.JSON(resp)
}
func (h *UserHandler) Patch(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	var req dto.UpdatePatchUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updated, err := h.Service.PatchUpdate(uint(id), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	resp := dto.UserResponse{
		ID:    updated.ID,
		Name:  updated.Name,
		Email: updated.Email,
		Role:  updated.Role,
	}

	return c.JSON(resp)
}
func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	if err := h.Service.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Deleted Data successfully",
		"data":    id,
	})
}

func (h *UserHandler) HardDelete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	if err := h.Service.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "status":  "success",
    "message": "Hard Deleted Data successfully",
    "data":    id,
	})


}


