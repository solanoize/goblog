package dtos

import (
	"time"
)

type UserResponseDTO struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	IsAdmin   bool      `json:"isAdmin"`
	IsStaff   bool      `json:"isStaff"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
