package dto

 type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"required,oneof=admin approver requester"` // e.g., "admin", "approver", "requester"
}

type UpdateUserRequest struct {
	ID       uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" validate:"omitempty,email"`
	Role  string `json:"role" validate:"omitempty,oneof=admin approver user"`
	Password string `json:"password"`

}
type UpdatePatchUserRequest struct {
	Name  		*string `json:"name"`
	Email 		*string `json:"email" validate:"omitempty,email"`
	Role  		*string `json:"role" validate:"omitempty,oneof=admin approver user"`
	Password 	*string `json:"password"`

}
type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	RoleLocal string 
	UserID     uint `json:"user_id"`
}