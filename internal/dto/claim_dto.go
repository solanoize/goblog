package dto

import "github.com/golang-jwt/jwt/v5"

type ClaimDTO struct {
	ID       uint
	Username string
	Email    string
	IsAdmin  bool `json:"isAdmin"`
	IsStaff  bool `json:"isStaff"`
	Active   bool `json:"active"`
	jwt.RegisteredClaims
}
