package users

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/solanoize/goblog/internal/auth"
	"github.com/solanoize/goblog/internal/utils"
)

type Middleware interface {
	IsAuthenticated(w http.ResponseWriter, r *http.Request, next http.Handler)
}

type middleware struct {
	Service Service
}

// IsAuthenticated implements [Middleware].
func (m *middleware) IsAuthenticated(w http.ResponseWriter, r *http.Request, next http.Handler) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		utils.RenderUnauthorized(w, "Authorization header required")
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		utils.RenderUnauthorized(w, "Authorization header format must be Bearer {token}")
		return
	}

	claimContract, err := m.Service.Auth().Parse(parts[1])
	if err != nil {
		utils.RenderUnauthorized(w, err.Error())
		return
	}

	claimContract, err = m.Service.GetForClaim(r.Context(), claimContract.UserID)

	ctx := context.WithValue(r.Context(), auth.CLAIM_CONTRACT, claimContract)
	next.ServeHTTP(w, r.WithContext(ctx))

}

func NewMiddleware(service Service, logger *log.Logger) Middleware {
	return &middleware{
		Service: service,
	}
}
