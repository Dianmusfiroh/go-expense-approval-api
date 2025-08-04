package services

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/dto"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/models"
)


type UserServiceInterface interface {
	Create(userDTO *dto.CreateUserRequest) (*dto.UserResponse, error)
	Update(id uint, userDTO *dto.UpdateUserRequest) (*dto.UserResponse, error)
	FindAll() ([]dto.UserResponse, error)
	FindByID(id uint) (*dto.UserResponse, error)
	Delete(id uint) error
	HardDelete(id uint) error
	PatchUpdate(id uint, req *dto.UpdatePatchUserRequest) (*models.User, error)
}

