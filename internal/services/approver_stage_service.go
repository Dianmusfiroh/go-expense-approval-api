package services

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/dto"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/models"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/repository"
)

type ApprovalStageService struct {
	Repo *repository.ApprovalStageRepository
}

func NewApprovalStageService(repo *repository.ApprovalStageRepository) *ApprovalStageService {
	return &ApprovalStageService{Repo: repo}
}

func (s *ApprovalStageService) Create(req *dto.CreateApprovalStageRequest) (*dto.ApprovalStageResponse, error) {
	stage := models.ApprovalStage{
		ApproverID: req.ApproverID,
		StageOrder: req.StageOrder,
	}

	if err := s.Repo.Create(&stage); err != nil {
		return nil, err
	}

	return &dto.ApprovalStageResponse{
		ID:         stage.ID,
		ApproverID: stage.ApproverID,
		StageOrder: stage.StageOrder,
	}, nil
}

func (s *ApprovalStageService) FindAll() ([]dto.ApprovalStageResponse, error) {
	stages, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []dto.ApprovalStageResponse
	for _, s := range stages {
		responses = append(responses, dto.ApprovalStageResponse{
			ID:         s.ID,
			ApproverID: s.ApproverID,
			StageOrder: s.StageOrder,
		})
	}
	return responses, nil
}

func (s *ApprovalStageService) FindByID(id uint) (*dto.ApprovalStageResponse, error) {
	stage, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.ApprovalStageResponse{
		ID:         stage.ID,
		ApproverID: stage.ApproverID,
		StageOrder: stage.StageOrder,
	}, nil
}

func (s *ApprovalStageService) Update(id uint, req *dto.UpdateApprovalStageRequest) (*dto.ApprovalStageResponse, error) {
	stage, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	stage.ApproverID = req.ApproverID
	stage.StageOrder = req.StageOrder

	if err := s.Repo.Update(stage); err != nil {
		return nil, err
	}

	return &dto.ApprovalStageResponse{
		ID:         stage.ID,
		ApproverID: stage.ApproverID,
		StageOrder: stage.StageOrder,
	}, nil
}

func (s *ApprovalStageService) Delete(id uint) error {
	return s.Repo.Delete(id)
}