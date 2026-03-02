package auth

import "github.com/golang-jwt/jwt/v5"

type ClaimKey string

const CLAIM_CONTRACT ClaimKey = "CLAIM_CONTRACT"

type ClaimsContract struct {
	UserID uint
	jwt.RegisteredClaims
}
