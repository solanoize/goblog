package usecases

import (
	"errors"
	"log"

	"github.com/solanoize/goblog/internal/dtos"
	"github.com/solanoize/goblog/internal/mappers"
	"github.com/solanoize/goblog/internal/models"
	"github.com/solanoize/goblog/internal/repositories"
)

type UserUseCase interface {
	Register(userRegisterDTO dtos.UserRegisterDTO) (dtos.UserResponseDTO, error)
	SignIn(userSignInDTO dtos.UserSignInDTO) (string, error)
	Me(id uint) (dtos.UserResponseDTO, error)
	GetAuthUseCase() AuthUseCase
	All(page int, limit int, search string) ([]dtos.UserResponseDTO, int64, error)
}

type userUseCase struct {
	Logger         *log.Logger
	UserRepository repositories.UserRepository
	AuthUseCase    AuthUseCase
}

// All implements [UserUseCase].
func (u *userUseCase) All(page int, limit int, search string) ([]dtos.UserResponseDTO, int64, error) {
	var err error
	var users []models.User
	var count int64
	var userResponseDTOS []dtos.UserResponseDTO
	var userResponseMapper mappers.UserResponseMapper

	users, count, err = u.UserRepository.FindAll(page, limit, search)
	userResponseDTOS = userResponseMapper.ToDTOS(users)
	if err != nil {
		return userResponseDTOS, count, errors.New("Terjadi kesalahan saat mengambil data user.")
	}

	return userResponseDTOS, count, err
}

// GetAuthUseCase implements [UserUseCase].
func (u *userUseCase) GetAuthUseCase() AuthUseCase {
	return u.AuthUseCase
}

// Me implements [UserUseCase].
func (u *userUseCase) Me(id uint) (dtos.UserResponseDTO, error) {
	var err error
	var user models.User
	var userResponseMapper mappers.UserResponseMapper
	var userResponseDTO dtos.UserResponseDTO

	user, err = u.UserRepository.FindById(id)
	userResponseDTO = userResponseMapper.ToDTO(user)
	if err != nil {
		u.Logger.Println(err)
		return userResponseDTO, errors.New("User tidak ditemukan.")
	}

	return userResponseDTO, err
}

// Register implements [UserUseCase].
func (u *userUseCase) Register(userRegisterDTO dtos.UserRegisterDTO) (dtos.UserResponseDTO, error) {
	var err error
	var userResponseDTO dtos.UserResponseDTO
	var userRegisterMapper mappers.UserRegisterMapper
	var userResponseMapper mappers.UserResponseMapper
	var user models.User

	userRegisterDTO.Password, err = u.AuthUseCase.PassowrdHash(userRegisterDTO.Password)
	if err != nil {
		return userResponseDTO, err
	}

	user, err = u.UserRepository.Create(userRegisterMapper.ToModel(userRegisterDTO))
	userResponseDTO = userResponseMapper.ToDTO(user)
	if err != nil {
		u.Logger.Println(err)
		return userResponseDTO, errors.New("User gagal dibuat, pastikan user belum pernah mendaftar sebelumnya.")
	}

	return userResponseDTO, nil
}

// SignIn implements [UserUseCase].
func (u *userUseCase) SignIn(userSignInDTO dtos.UserSignInDTO) (string, error) {
	var err error
	var user models.User
	var token string

	user, err = u.UserRepository.FindByEmail(userSignInDTO.Email)
	if err != nil {
		u.Logger.Println(err)
		return token, errors.New("Email belum terdaftar.")
	}

	err = u.AuthUseCase.ComparePassword(userSignInDTO.Password, user.Password)
	if err != nil {
		return token, err
	}

	token, err = u.AuthUseCase.GenerateToken(user)
	if err != nil {
		return token, err
	}

	return token, nil
}

func NewUserUseCase(logger *log.Logger, userRepository repositories.UserRepository) UserUseCase {
	return &userUseCase{
		Logger:         logger,
		UserRepository: userRepository,
		AuthUseCase:    NewAuthUseCase(logger),
	}
}
