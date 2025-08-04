package services

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/dto"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/models"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/repository"
)

type ApproverService struct {
	Repo *repository.ApproverRepository
}

func NewApproverService(repo *repository.ApproverRepository) *ApproverService {
	return &ApproverService{Repo: repo}
}

func (s *ApproverService) Create(req *dto.CreateApproverRequest) (*dto.ApproverResponse, error) {
	approver := models.Approver{Name: req.Name}
	if err := s.Repo.Create(&approver); err != nil {
		return nil, err
	}
	return &dto.ApproverResponse{ID: approver.ID, Name: approver.Name, Email: approver.Email}, nil
}

func (s *ApproverService) FindAll() ([]dto.ApproverResponse, error) {
	approvers, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}
	var responses []dto.ApproverResponse
	for _, a := range approvers {
		responses = append(responses, dto.ApproverResponse{ID: a.ID, Name: a.Name})
	}
	return responses, nil
}

func (s *ApproverService) FindByID(id uint) (*dto.ApproverResponse, error) {
	a, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.ApproverResponse{ID: a.ID, Name: a.Name, Email: a.Email}, nil
}

func (s *ApproverService) Update(id uint, req *dto.UpdateApproverRequest) (*dto.ApproverResponse, error) {
	a, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	a.Name = req.Name
	if err := s.Repo.Update(a); err != nil {
		return nil, err
	}
	return &dto.ApproverResponse{ID: a.ID, Name: a.Name, Email: a.Email }, nil
}

func (s *ApproverService) Delete(id uint) error {
	return s.Repo.Delete(id)
}
