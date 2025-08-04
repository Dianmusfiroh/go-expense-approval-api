package seeder

import (
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/models"
	"gorm.io/gorm"
)

var isSeeded bool

func Seed(db *gorm.DB) {
	if isSeeded {
		return
	}

	isSeeded = true

	// Seed statuses
	statuses := []models.Status{
		{Name: "menunggu persetujuan"},
		{Name: "disetujui"},
		{Name: "ditolak"},
	}
	db.Create(&statuses)

	// Seed users
	admin := models.User{
		Name:     "Admin",
		Email:    "admin@example.com",
		Password: "admin123", // sebaiknya di-hash (bcrypt)
		Role:     "admin",
	}
	db.Create(&admin)

	requester := models.User{
		Name:     "Requester",
		Email:    "requester@example.com",
		Password: "requester123",
		Role:     "requester",
	}
	db.Create(&requester)

	// Seed approvers
	approver1 := models.Approver{Name: "Manager A", Email: "managerA@example.com"}
	approver2 := models.Approver{Name: "Manager B", Email: "managerB@example.com"}
	db.Create(&approver1)
	db.Create(&approver2)

	// Seed approval stages (urutan penting)
	stages := []models.ApprovalStage{
		{ApproverID: approver1.ID, StageOrder: 1},
		{ApproverID: approver2.ID, StageOrder: 2},
	}
	db.Create(&stages)

	// Seed one expense from requester
	expense := models.Expense{
		UserID:      requester.ID,
		Amount:      150000,
		Description: "Biaya transportasi dinas",
		StatusID:    statuses[0].ID, // menunggu persetujuan
	}
	db.Create(&expense)

	// Seed approval record for first approver (pending)
	approval := models.Approval{
		ExpenseID:  expense.ID,
		ApproverID: approver1.ID,
		StatusID:   statuses[0].ID, // menunggu persetujuan
		ApprovedAt: nil,            // belum disetujui
	}
	db.Create(&approval)
}

