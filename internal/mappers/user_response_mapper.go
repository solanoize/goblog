package mappers

import (
	"github.com/solanoize/goblog/internal/dtos"
	"github.com/solanoize/goblog/internal/models"
)

type UserResponseMapper interface {
	ToDTO(user models.User) dtos.UserResponseDTO
	ToDTOS(users []models.User) []dtos.UserResponseDTO
	ToModel(userResponseDTO dtos.UserResponseDTO) models.User
}

type userResponseMapper struct {
}

// ToModel implements [UserResponseMapper].
func (m *userResponseMapper) ToModel(userResponseDTO dtos.UserResponseDTO) models.User {
	return models.User{
		Username: userResponseDTO.Username,
		Email:    userResponseDTO.Email,
		IsAdmin:  userResponseDTO.IsAdmin,
		IsStaff:  userResponseDTO.IsStaff,
		Active:   userResponseDTO.Active,
	}
}

func (m *userResponseMapper) ToDTO(user models.User) dtos.UserResponseDTO {
	return dtos.UserResponseDTO{
		Username:  user.Username,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		IsStaff:   user.IsStaff,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (m *userResponseMapper) ToDTOS(users []models.User) []dtos.UserResponseDTO {
	var userResponseDTOS []dtos.UserResponseDTO = make([]dtos.UserResponseDTO, 0, len(users))
	var user models.User

	for _, user = range users {
		userResponseDTOS = append(userResponseDTOS, m.ToDTO(user))
	}

	return userResponseDTOS
}

func NewUserResponseMapper() UserResponseMapper {
	return &userResponseMapper{}
}
