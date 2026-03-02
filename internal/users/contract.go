package users

type UserResponseContract struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"isAdmin"`
	IsStaff  bool   `json:"isStaff"`
	Active   bool   `json:"active"`
}

type TokenResponseContract struct {
	Token string `json:"token"`
}

type SignInRequestContract struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type SignUpRequestContract struct {
	Username string `json:"username" validate:"required,min=5"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
