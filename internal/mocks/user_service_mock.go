package mocks

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/dto"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/models"
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

// Create implements services.UserServiceInterface.
func (m *UserServiceMock) Create(userDTO *dto.CreateUserRequest) (*dto.UserResponse, error) {
	panic("unimplemented")
}

// Delete implements services.UserServiceInterface.
func (m *UserServiceMock) Delete(id uint) error {
	panic("unimplemented")
}

// FindByID implements services.UserServiceInterface.
func (m *UserServiceMock) FindByID(id uint) (*dto.UserResponse, error) {
	panic("unimplemented")
}

// HardDelete implements services.UserServiceInterface.
func (m *UserServiceMock) HardDelete(id uint) error {
	panic("unimplemented")
}

// PatchUpdate implements services.UserServiceInterface.
func (m *UserServiceMock) PatchUpdate(id uint, req *dto.UpdatePatchUserRequest) (*models.User, error) {
	panic("unimplemented")
}

// Update implements services.UserServiceInterface.
func (m *UserServiceMock) Update(id uint, userDTO *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	panic("unimplemented")
}

func (m *UserServiceMock) FindAll() ([]dto.UserResponse, error) {
	args := m.Called()
	return args.Get(0).([]dto.UserResponse), args.Error(1)
}
