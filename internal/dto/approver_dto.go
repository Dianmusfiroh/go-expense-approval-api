package dto

type CreateApproverRequest struct {
	Name     string `json:"name"`
}
type UpdateApproverRequest struct {
	Name     string `json:"name"`
}
type ApproverResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}