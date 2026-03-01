package mappers

import (
	"github.com/solanoize/goblog/internal/dtos"
	"github.com/solanoize/goblog/internal/models"
)

type UserSignInMapper struct{}

func (u *UserSignInMapper) ToDTO(user models.User) dtos.UserSignInDTO {
	return dtos.UserSignInDTO{
		Email:    user.Email,
		Password: user.Password,
	}
}
