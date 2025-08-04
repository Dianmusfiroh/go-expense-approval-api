package services

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/dto"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/models"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository

}
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) Create(userDTO *dto.CreateUserRequest) (*dto.UserResponse, error) {
	user := models.User{
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Password: userDTO.Password, // idealnya hash dulu
		Role:     userDTO.Role,
	}

	if err := s.Repo.Create(&user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (s *UserService) Update(id uint, userDTO *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = userDTO.Name
	user.Email = userDTO.Email
	user.Password = userDTO.Password
	user.Role = userDTO.Role

	if err := s.Repo.Update(user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}


func (s *UserService) FindAll() ([]dto.UserResponse, error) {
	users, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}

	var result []dto.UserResponse
	for _, u := range users {
		result = append(result, dto.UserResponse{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
			Role:  u.Role,
		})
	}

	return result, nil
}

func (s *UserService) FindByID(id uint) (*dto.UserResponse, error) {
	user, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (s *UserService) Delete(id uint) error {
	return s.Repo.Delete(id)
}
func (s *UserService) HardDelete(id uint) error {
	return s.Repo.Delete(id)
}
func (s *UserService) PatchUpdate(id uint, req *dto.UpdatePatchUserRequest) (*models.User, error) {
	user, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Password != nil {
		user.Password = *req.Password // encrypt if needed
	}
	if req.Role != nil {
		user.Role = *req.Role
	}

	if err := s.Repo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}
