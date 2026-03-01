package mappers

import (
	"github.com/solanoize/goblog/internal/dtos"
	"github.com/solanoize/goblog/internal/models"
)

type UserResponseMapper struct {
}

func (m *UserResponseMapper) ToDTO(user models.User) dtos.UserResponseDTO {
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

func (m *UserResponseMapper) ToDTOS(users []models.User) []dtos.UserResponseDTO {
	var userResponseDTOS []dtos.UserResponseDTO = make([]dtos.UserResponseDTO, 0, len(users))
	var user models.User

	for _, user = range users {
		userResponseDTOS = append(userResponseDTOS, m.ToDTO(user))
	}

	return userResponseDTOS
}
