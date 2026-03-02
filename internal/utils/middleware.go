package utils

import "net/http"

func MiddlewareAdapter(fn func(http.ResponseWriter, *http.Request, http.Handler)) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fn(w, r, next)
		})
	}
}
