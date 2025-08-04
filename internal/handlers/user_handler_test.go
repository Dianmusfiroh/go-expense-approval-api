package handlers_test

import (
	"net/http/httptest"
	"testing"

	"github.com/Dianmusfiroh/go-expense-approval-api/internal/dto"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/handlers"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	app := fiber.New()

	mockService := new(mocks.UserServiceMock)
	handler := handlers.NewUserHandler(mockService)

	expectedUsers := []dto.UserResponse{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Role: "admin"},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Role: "user"},
	}

	mockService.On("FindAll").Return(expectedUsers, nil)

	app.Get("/users", handler.GetAll)

	req := httptest.NewRequest("GET", "/users", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	mockService.AssertExpectations(t)
}
