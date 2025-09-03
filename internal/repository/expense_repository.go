package repository

import (
	"gorm.io/gorm"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/models"
)
type ExpenseRepository interface {
	Create(expense *models.Expense) error
	FindAll() ([]models.Expense, error)
}

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepository{db}
}

func (r *expenseRepository) Create(expense *models.Expense) error {
	return r.db.Create(expense).Error
}

func (r *expenseRepository) FindAll() ([]models.Expense, error) {
	var expenses []models.Expense
	err := r.db.Preload("User").Preload("Status").Find(&expenses).Error
	return expenses, err
}