package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/solanoize/goblog/internal/dtos"
	"github.com/solanoize/goblog/internal/usecases"
)

type AuthMiddleware interface {
	IsAuthenticated() func(http.Handler) http.Handler
	IsAdminOnly() func(http.Handler) http.Handler
	IsActiveUser() func(http.Handler) http.Handler
}

type authMiddleware struct {
	Logger      *log.Logger
	UserUseCase usecases.UserUseCase
}

// IsActiveUser implements [AuthMiddleware].
func (a *authMiddleware) IsActiveUser() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dto, ok := r.Context().Value("dto").(dtos.UserResponseDTO)

			if !ok {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(map[string]string{"detail": "Invalid context key"})
				return
			}

			if !dto.Active {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(map[string]string{"detail": "Akses ditolak karena akun tidak aktif"})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// IsAdminOnly implements [AuthMiddleware].
func (a *authMiddleware) IsAdminOnly() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var userResponseDTO dtos.UserResponseDTO
			var ok bool

			userResponseDTO, ok = r.Context().Value("dto").(dtos.UserResponseDTO)

			if !ok {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(map[string]string{"detail": "Invalid context key"})
				return
			}

			if !userResponseDTO.Active {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(map[string]string{"detail": "Akses ditolak karena akun tidak aktif"})
				return
			}

			if !userResponseDTO.IsAdmin {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(map[string]string{"detail": "Akses ditolak karena bukan admin"})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// Protected implements [AuthMiddleware].
func (a *authMiddleware) IsAuthenticated() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var parts []string
			var userResponseDTO dtos.UserResponseDTO
			var authHeader string
			var err error
			var tokenString string
			var userID uint

			authHeader = r.Header.Get("Authorization")
			if authHeader == "" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"detail": "Authorization header required"})
				return
			}

			parts = strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"detail": "Authorization header format must be Bearer {token}"})
				return
			}

			tokenString = parts[1]
			userID, err = a.UserUseCase.GetAuthUseCase().ParseToken(tokenString)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"detail": err.Error()})
				return
			}

			if userID == 0 {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"detail": "User tidak ditemukan."})
				return
			}

			userResponseDTO, err = a.UserUseCase.Me(userID)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"detail": "User tidak ditemukan."})
				return
			}

			var ctx context.Context
			ctx = context.WithValue(r.Context(), "dto", userResponseDTO)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

func NewAuthMiddleware(logger *log.Logger, userUseCase usecases.UserUseCase) AuthMiddleware {
	return &authMiddleware{Logger: logger, UserUseCase: userUseCase}
}
