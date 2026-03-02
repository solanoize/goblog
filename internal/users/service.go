package users

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/solanoize/goblog/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	SignUp(context context.Context, requestContract SignUpRequestContract) (UserResponseContract, error)
	SignIn(context context.Context, requestContract SignInRequestContract) (TokenResponseContract, error)
	GetByID(context context.Context, id uint) (UserResponseContract, error)
	GetForClaim(context context.Context, id uint) (auth.ClaimsContract, error)
	Auth() auth.Service
}

type service struct {
	Logger      *log.Logger
	Repository  Repository
	AuthService auth.Service
}

// GetForClaim implements [Service].
func (s *service) GetForClaim(context context.Context, id uint) (auth.ClaimsContract, error) {
	user, err := s.Repository.FindByID(context, id)
	if err != nil {
		s.Logger.Println(err)
		return auth.ClaimsContract{}, errors.New("Tidak dapat mengambil detail user")
	}

	return auth.ClaimsContract{
		UserID: user.ID,
	}, err
}

// Authentication implements [Service].
func (s *service) Auth() auth.Service {
	return s.AuthService
}

// GetById implements [Service].
func (s *service) GetByID(context context.Context, id uint) (UserResponseContract, error) {
	user, err := s.Repository.FindByID(context, id)
	if err != nil {
		s.Logger.Println(err)
		return UserResponseContract{}, errors.New("Tidak dapat mengambil detail user")
	}

	return UserResponseContract{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		IsAdmin:  user.IsAdmin,
		IsStaff:  user.IsStaff,
		Active:   user.Active,
	}, err
}

// SignIn implements [Service].
func (s *service) SignIn(context context.Context, requestContract SignInRequestContract) (TokenResponseContract, error) {
	user, err := s.Repository.FindByEmail(context, requestContract.Email)

	if err != nil {
		s.Logger.Println(err)
		return TokenResponseContract{}, errors.New("Email tidak terdaftar")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestContract.Password))
	if err != nil {
		s.Logger.Println(err)
		return TokenResponseContract{}, errors.New("Password tidak cocok.")
	}

	claims := auth.ClaimsContract{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(120 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		s.Logger.Println(err)
		return TokenResponseContract{}, errors.New("Kredensial token gagal dibuat")
	}

	return TokenResponseContract{Token: tokenString}, nil
}

// SignUp implements [Service].
func (s *service) SignUp(context context.Context, requestContract SignUpRequestContract) (UserResponseContract, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestContract.Password), bcrypt.DefaultCost)
	if err != nil {
		s.Logger.Println(err)
		return UserResponseContract{}, errors.New("Terjadi kesalahan saat hashing password")
	}

	user, err := s.Repository.Create(context, User{
		Username: requestContract.Username,
		Email:    requestContract.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		s.Logger.Println(err)
		return UserResponseContract{}, errors.New("Terjadi kesalahan saat membuat user")
	}

	return UserResponseContract{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		IsAdmin:  user.IsAdmin,
		IsStaff:  user.IsStaff,
		Active:   user.Active,
	}, err
}

func NewService(repository Repository, authService auth.Service, logger *log.Logger) Service {
	return &service{
		Logger:      logger,
		Repository:  repository,
		AuthService: authService,
	}
}
