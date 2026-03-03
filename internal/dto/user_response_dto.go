package dto

type UserResponseDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"isAdmin"`
	IsStaff  bool   `json:"isStaff"`
	Active   bool   `json:"active"`
}
