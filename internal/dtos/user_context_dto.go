package dtos

type UserContextDTO struct {
	ID       uint
	Username string
	Email    string
	IsAdmin  bool
	IsStaff  bool
	Active   bool
}
