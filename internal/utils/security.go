package utils

import (
	"errors"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/solanoize/goblog/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

func ParseHandlerSecurity(t *jwt.Token) (any, error) {
	_, ok := t.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("Unexpected signing method algorithm.")
	}

	return []byte(os.Getenv("JWT_SECRET")), nil
}

func HashPasswordSecurity(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func ParseTokenSecurity(tokenString string) (dto.ClaimDTO, error) {
	claims := dto.ClaimDTO{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, ParseHandlerSecurity)
	if err != nil || !token.Valid {
		return dto.ClaimDTO{}, err
	}

	return claims, nil
}

func ComparePassword(incomingPassword string, originalPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(originalPassword), []byte(incomingPassword))
}

func GetClaim(r *http.Request) (dto.ClaimDTO, error) {
	claims, ok := r.Context().Value(CLAIM_CTYPE).(dto.ClaimDTO)
	if !ok {
		return dto.ClaimDTO{}, errors.New("Invalid session, please try again to signin")
	}

	return claims, nil
}
