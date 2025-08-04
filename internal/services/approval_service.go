package services

import (
	"errors"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/dto"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/repository"
)

type ApprovalService struct {
	Repo *repository.ApprovalRepository
}

func NewApprovalService(repo *repository.ApprovalRepository) *ApprovalService {
	return &ApprovalService{Repo: repo}
}

func (s *ApprovalService) ApproveExpense(input *dto.ApproveExpenseRequest) error {
	nextApproval, err := s.Repo.FindNextApproval(input.ExpenseID)
	if err != nil {
		return errors.New("tidak ada approval yang bisa dilakukan")
	}

	if nextApproval.ApproverID != input.ApproverID {
		return errors.New("anda belum bisa menyetujui pengeluaran ini")
	}

	if err := s.Repo.Approve(nextApproval); err != nil {
		return err
	}

	allApproved, err := s.Repo.AllStagesApproved(input.ExpenseID)
	if err != nil {
		return err
	}

	if allApproved {
		return s.Repo.UpdateExpenseStatus(input.ExpenseID, 2) // disetujui
	}

	return nil
}