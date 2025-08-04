package services

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/models"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/repository"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/dto"
)

type ExpenseService interface {
	Create(dto.CreateExpenseRequest) (*models.Expense, error)
	GetAll() ([]models.Expense, error)
}

type expenseService struct {
	repo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) ExpenseService {
	return &expenseService{repo}
}

func (s *expenseService) Create(req dto.CreateExpenseRequest) (*models.Expense, error) {
	expense := models.Expense{
		UserID:      req.UserID,
		Amount:      req.Amount,
		Description: req.Description,
		StatusID:    1, // diasumsikan 1 = "menunggu persetujuan"
	}
	if err := s.repo.Create(&expense); err != nil {
		return nil, err
	}
	return &expense, nil
}

func (s *expenseService) GetAll() ([]models.Expense, error) {
	return s.repo.FindAll()
}