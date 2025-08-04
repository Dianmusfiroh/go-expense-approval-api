package models

import (
	"time"
	"gorm.io/gorm"
)

type Approval struct {
	gorm.Model
	ExpenseID  uint
	Expense    Expense
	ApproverID uint
	Approver   Approver
	StatusID   uint
	Status     Status
	ApprovedAt *time.Time
}
