package utils

import (
	"encoding/json"
	"net/http"
)

func JSONRender(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

func ErrorRender(w http.ResponseWriter, status int, message any) {
	JSONRender(w, status, map[string]any{"detail": message})
}

func OKRender(w http.ResponseWriter, payload any) {
	JSONRender(w, http.StatusOK, payload)
}

func CreatedRender(w http.ResponseWriter, payload any) {
	JSONRender(w, http.StatusCreated, payload)
}

func NoContentRender(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func BadRequestRender(w http.ResponseWriter, message any) {
	ErrorRender(w, http.StatusBadRequest, message)
}

func UnauthorizedRender(w http.ResponseWriter, message any) {
	ErrorRender(w, http.StatusUnauthorized, message)
}

func ForbiddenRender(w http.ResponseWriter, message any) {
	ErrorRender(w, http.StatusForbidden, message)
}

func NotFoundRender(w http.ResponseWriter, message any) {
	ErrorRender(w, http.StatusNotFound, message)
}

func ConflictRender(w http.ResponseWriter, message any) {
	ErrorRender(w, http.StatusConflict, message)
}

func InternalServerErrorRender(w http.ResponseWriter, message any) {
	ErrorRender(w, http.StatusInternalServerError, message)
}
