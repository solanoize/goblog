package auth

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Hash(password string) (string, error)
	Compare(incomingPassword string, originalPassword string) error
	Parse(tokenString string) (ClaimsContract, error)
	GenerateToken(id uint) (string, error)
	ParseHandler(t *jwt.Token) (any, error)
	GetClaim(r *http.Request) (ClaimsContract, error)
}

type service struct {
	Logger *log.Logger
}

// GetClaim implements [Service].
func (s *service) GetClaim(r *http.Request) (ClaimsContract, error) {
	claims, ok := r.Context().Value(CLAIM_CONTRACT).(ClaimsContract)
	if !ok {
		return claims, errors.New("Sesi tidak valid, silakan login ulang")
	}
	return claims, nil
}

// ParseHandler implements [Service].
func (s *service) ParseHandler(t *jwt.Token) (any, error) {
	_, ok := t.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		s.Logger.Printf("Unexpected signing method: %v", t.Header["alg"])
		return nil, errors.New("Unexpected signing method algorithm.")
	}

	return []byte(os.Getenv("JWT_SECRET")), nil
}

// Compare implements [Service].
func (s *service) Compare(incomingPassword string, originalPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(originalPassword), []byte(incomingPassword))
	if err != nil {
		s.Logger.Println(err)
		return errors.New("Password not same")
	}

	return nil
}

// GenerateToken implements [Service].
func (s *service) GenerateToken(id uint) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(120 * time.Hour).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		s.Logger.Println(err)
		return "", errors.New("Kredensial token gagal dibuat")
	}

	return token, nil
}

// Hash implements [Service].
func (s *service) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.Logger.Println(err)
		return "", errors.New("Terjadi kesalahan saat menghasilkan password")
	}

	return string(hashedPassword), nil
}

// Parse implements [Service].
func (s *service) Parse(tokenString string) (ClaimsContract, error) {
	claims := &ClaimsContract{}

	token, err := jwt.ParseWithClaims(tokenString, claims, s.ParseHandler)
	if err != nil || !token.Valid {
		return ClaimsContract{}, errors.New("invalid or expired token")
	}

	return *claims, nil
}

func NewService(logger *log.Logger) Service {
	return &service{Logger: logger}
}
