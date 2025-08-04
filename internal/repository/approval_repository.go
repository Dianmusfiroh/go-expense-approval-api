package repository

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/models"
	"gorm.io/gorm"
)
type ApprovalRepository struct {
	DB *gorm.DB
}

func NewApprovalRepository(db *gorm.DB) *ApprovalRepository {
	return &ApprovalRepository{DB: db}
}

func (r *ApprovalRepository) FindNextApproval(expenseID uint) (*models.Approval, error) {
	var approval models.Approval
	err := r.DB.
		Joins("JOIN approval_stages ON approvals.approver_id = approval_stages.approver_id").
		Where("expense_id = ? AND approvals.status_id = ?", expenseID, 1). // 1: menunggu persetujuan
		Order("approval_stages.id ASC").
		First(&approval).Error

	return &approval, err
}

func (r *ApprovalRepository) Approve(approval *models.Approval) error {
	return r.DB.Model(approval).Update("status_id", 2).Error // 2: disetujui
}

func (r *ApprovalRepository) AllStagesApproved(expenseID uint) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Approval{}).
		Where("expense_id = ? AND status_id = 1", expenseID). // masih ada yg menunggu
		Count(&count).Error
	return count == 0, err
}

func (r *ApprovalRepository) UpdateExpenseStatus(expenseID uint, statusID uint) error {
	return r.DB.Model(&models.Expense{}).Where("id = ?", expenseID).Update("status_id", statusID).Error
}