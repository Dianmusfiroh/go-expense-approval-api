package repository

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/models"
	"gorm.io/gorm"
)

type ApproverRepository struct {
	DB *gorm.DB
}

func NewApproverRepository(db *gorm.DB) *ApproverRepository {
	return &ApproverRepository{DB: db}
}

func (r *ApproverRepository) Create(approver *models.Approver) error {
	return r.DB.Create(approver).Error
}

func (r *ApproverRepository) FindAll() ([]models.Approver, error) {
	var approvers []models.Approver
	err := r.DB.Find(&approvers).Error
	return approvers, err
}

func (r *ApproverRepository) FindByID(id uint) (*models.Approver, error) {
	var approver models.Approver
	err := r.DB.First(&approver, id).Error
	return &approver, err
}

func (r *ApproverRepository) Update(approver *models.Approver) error {
	return r.DB.Save(approver).Error
}

func (r *ApproverRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Approver{}, id).Error
}
