package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	UserID      uint
	User        User
	Amount      float64
	Description string
	StatusID    uint
	Status      Status
	Code        string
	Approvals   []Approval `gorm:"foreignKey:ExpenseID"`
}
