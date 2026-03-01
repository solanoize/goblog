package mappers

import (
	"github.com/solanoize/goblog/internal/dtos"
	"github.com/solanoize/goblog/internal/models"
)

type UserRegisterMapper struct{}

func (u *UserRegisterMapper) ToModel(userRegisterDTO dtos.UserRegisterDTO) models.User {
	return models.User{
		Email:    userRegisterDTO.Email,
		Username: userRegisterDTO.Username,
		Password: userRegisterDTO.Password,
	}
}
