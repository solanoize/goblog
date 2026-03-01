package usecases

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/solanoize/goblog/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	PassowrdHash(password string) (string, error)
	ComparePassword(incomingPassword string, password string) error
	ParseToken(tokenString string) (uint, error)
	GenerateToken(user models.User) (string, error)
}

type authUseCase struct {
	Logger *log.Logger
}

// GenerateToken implements [AuthUseCase].
func (a *authUseCase) GenerateToken(user models.User) (string, error) {
	var secret string = os.Getenv("JWT_SECRET")
	var err error
	var jwtToken *jwt.Token
	var token string

	var claims jwt.Claims = jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(120 * time.Hour).Unix(),
	}

	jwtToken = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(secret))
	if err != nil {
		a.Logger.Println(err)
		return "", errors.New("Kredensial token gagal dibuat.")
	}

	return token, nil
}

// ParseToken implements [AuthUseCase].
func (a *authUseCase) ParseToken(tokenString string) (uint, error) {
	var err error
	var parseFunc func(t *jwt.Token) (any, error)
	var token *jwt.Token
	var claims jwt.MapClaims
	var ok bool
	var uidAny interface{}
	var uid float64

	parseFunc = func(t *jwt.Token) (any, error) {
		var ok bool
		_, ok = t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			a.Logger.Printf("Unexpected signing method: %v", t.Header["alg"])
			return nil, errors.New("Unexpected signing method algorithm.")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil

	}

	token, err = jwt.Parse(tokenString, parseFunc)
	if err != nil || !token.Valid {
		a.Logger.Println(err)
		return 0, errors.New("Invalid token.")
	}

	claims, ok = token.Claims.(jwt.MapClaims)

	if !ok {
		a.Logger.Println("Invalid token claims.")
		return 0, errors.New("Invalid token claims.")
	}

	uidAny, ok = claims["userID"]

	if !ok {
		a.Logger.Println("user id tidak ada di claims.")
		return 0, errors.New("Invalid user id")
	}

	uid, ok = uidAny.(float64)
	a.Logger.Println(uint(uid))
	if !ok {
		a.Logger.Println("Invalid user id type.")
		return 0, errors.New("Invalid user id type.")
	}

	return uint(uid), nil

}

// ComparePassword implements [AuthUseCase].
func (a *authUseCase) ComparePassword(incomingPassword string, password string) error {
	var err error

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(incomingPassword))
	if err != nil {
		a.Logger.Println(err)
		return errors.New("Password tidak cocok.")
	}

	return nil
}

// PassowrdHash implements [AuthUseCase].
func (a *authUseCase) PassowrdHash(password string) (string, error) {
	var err error
	var hashedPassword []byte

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		a.Logger.Println(err)
		return "", errors.New("Terjadi kesalahan saat menggenerate password.")
	}

	return string(hashedPassword), nil
}

func NewAuthUseCase(logger *log.Logger) AuthUseCase {
	return &authUseCase{
		Logger: logger,
	}
}
