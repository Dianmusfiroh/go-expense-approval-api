package models

import "gorm.io/gorm"

type ApprovalStage struct {
	gorm.Model
	ApproverID uint 
	Approver   Approver `gorm:"foreignKey:ApproverID;references:ID"`
	StageOrder int
}
