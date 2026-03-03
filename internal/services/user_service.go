package services

import (
	"context"
	"errors"
	"log"

	"github.com/solanoize/goblog/internal/dto"
	"github.com/solanoize/goblog/internal/models"
	"github.com/solanoize/goblog/internal/repositories"
	"github.com/solanoize/goblog/internal/utils"
)

type UserService interface {
	SignUp(ctx context.Context, requestDTO dto.SignUpRequestDTO) (dto.UserResponseDTO, error)
}

type userService struct {
	Logger         *log.Logger
	UserRepository repositories.UserRepository
}

// SignUp implements [UserService].
func (u *userService) SignUp(ctx context.Context, requestDTO dto.SignUpRequestDTO) (dto.UserResponseDTO, error) {
	hashedPassword, err := utils.HashPasswordSecurity(requestDTO.Password)
	if err != nil {
		u.Logger.Println(err)
		return dto.UserResponseDTO{}, errors.New("Invalid password hashing")
	}

	user, err := u.UserRepository.Create(ctx, models.User{
		Username: requestDTO.Username,
		Email:    requestDTO.Email,
		Password: hashedPassword,
	})
	if err != nil {
		u.Logger.Println(err)
		return dto.UserResponseDTO{}, errors.New("Failed to create user data")
	}

	return dto.UserResponseDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		IsAdmin:  user.IsAdmin,
		IsStaff:  user.IsStaff,
		Active:   user.Active,
	}, nil
}

func NewUserService(logger *log.Logger, userRepository repositories.UserRepository) UserService {
	return &userService{
		Logger:         logger,
		UserRepository: userRepository,
	}
}
