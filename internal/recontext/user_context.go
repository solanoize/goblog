package recontext

import (
	"context"

	"github.com/solanoize/goblog/internal/dtos"
)

type key string

const userKey key = "user_context_dto"

type UserContext interface {
	GetUserContext(ctx context.Context) dtos.UserContextDTO
}

type userContext struct{}

// GetUser implements [UserContext].
func (u *userContext) GetUserContext(ctx context.Context) dtos.UserContextDTO {
	var userContextDTO dtos.UserContextDTO
	userContextDTO, _ = ctx.Value("user_context_dto").(dtos.UserContextDTO)

	return userContextDTO

}

func NewUserContext() UserContext {
	return &userContext{}
}
