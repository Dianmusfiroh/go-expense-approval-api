package dto

type ApproveExpenseRequest struct {
	ApproverID uint `json:"approver_id"`
	ExpenseID  uint `json:"expense_id"`
}
