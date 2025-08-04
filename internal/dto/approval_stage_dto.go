package dto

type CreateApprovalStageRequest struct {
	ApproverID uint `json:"approver_id"`
	StageOrder int  `json:"stage_order"`
}

type UpdateApprovalStageRequest struct {
	ApproverID uint `json:"approver_id"`
	StageOrder int  `json:"stage_order"`
}

type ApprovalStageResponse struct {
	ID         uint `json:"id"`
	ApproverID uint `json:"approver_id"`
	StageOrder int  `json:"stage_order"`
}