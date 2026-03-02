package auth

import "net/http"

type Middleware interface {
	IsAuthenticated(w http.ResponseWriter, r *http.Request, next http.Handler)
}

type middleware struct {
	Service Service
}
