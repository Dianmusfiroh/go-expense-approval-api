package models

import (
	"gorm.io/gorm"
)


type Approver struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100)"`
	Email string `gorm:"uniqueIndex;type:varchar(100)"`
	ApprovalStages []ApprovalStage `gorm:"foreignKey:ApproverID"`
	Approvals      []Approval      `gorm:"foreignKey:ApproverID"`
}