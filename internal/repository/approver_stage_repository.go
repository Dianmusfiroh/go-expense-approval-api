package repository

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/models"
	"gorm.io/gorm"
)
type ApprovalStageRepository struct {
	DB *gorm.DB
}

func NewApprovalStageRepository(db *gorm.DB) *ApprovalStageRepository {
	return &ApprovalStageRepository{DB: db}
}

func (r *ApprovalStageRepository) Create(stage *models.ApprovalStage) error {
	return r.DB.Create(stage).Error
}

func (r *ApprovalStageRepository) FindAll() ([]models.ApprovalStage, error) {
	var stages []models.ApprovalStage
	err := r.DB.Preload("Approver").Find(&stages).Error
	return stages, err
}

func (r *ApprovalStageRepository) FindByID(id uint) (*models.ApprovalStage, error) {
	var stage models.ApprovalStage
	err := r.DB.Preload("Approver").First(&stage, id).Error
	return &stage, err
}

func (r *ApprovalStageRepository) Update(stage *models.ApprovalStage) error {
	return r.DB.Save(stage).Error
}

func (r *ApprovalStageRepository) Delete(id uint) error {
	return r.DB.Delete(&models.ApprovalStage{}, id).Error
}