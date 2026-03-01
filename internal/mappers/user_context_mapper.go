package mappers

import (
	"github.com/solanoize/goblog/internal/dtos"
	"github.com/solanoize/goblog/internal/models"
)

type UserContextMapper interface {
	ToDTO(user models.User) dtos.UserContextDTO
}

type userContextMapper struct{}

// ToDTO implements [UserContextMapper].
func (u *userContextMapper) ToDTO(user models.User) dtos.UserContextDTO {
	return dtos.UserContextDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		IsAdmin:  user.IsAdmin,
		IsStaff:  user.IsStaff,
		Active:   user.Active,
	}
}

func NewUserContextMapper() UserContextMapper {
	return &userContextMapper{}
}
