package models

import (
	"gorm.io/gorm"
)
// User represents a user in the system
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"` // e.g., "admin", "approver", "requester"
	Expenses []Expense  `gorm:"foreignKey:UserID"`

}