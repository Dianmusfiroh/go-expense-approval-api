package dto

type CreateExpenseRequest struct {
	Amount      float64 `json:"amount" validate:"required,gt=0"`
	Description string  `json:"description" validate:"required"`
	UserID      uint    `json:"user_id" validate:"required"`
}

type ExpenseResponse struct {
	ID          uint    `json:"id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	Code        string  `json:"code"`
}
