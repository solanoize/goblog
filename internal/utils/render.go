package utils

import (
	"encoding/json"
	"net/http"
)

func RenderJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

func RenderError(w http.ResponseWriter, status int, message any) {
	RenderJSON(w, status, map[string]any{"detail": message})
}

func RenderOK(w http.ResponseWriter, payload any) {
	RenderJSON(w, http.StatusOK, payload)
}

func RenderCreated(w http.ResponseWriter, payload any) {
	RenderJSON(w, http.StatusCreated, payload)
}

func RenderNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func RenderBadRequest(w http.ResponseWriter, message any) {
	RenderError(w, http.StatusBadRequest, message)
}

func RenderUnauthorized(w http.ResponseWriter, message any) {
	RenderError(w, http.StatusUnauthorized, message)
}

func RenderForbidden(w http.ResponseWriter, message any) {
	RenderError(w, http.StatusForbidden, message)
}

func RenderNotFound(w http.ResponseWriter, message any) {
	RenderError(w, http.StatusNotFound, message)
}

func RenderConflict(w http.ResponseWriter, message any) {
	RenderError(w, http.StatusConflict, message)
}

func RenderInternalServerError(w http.ResponseWriter, message any) {
	RenderError(w, http.StatusInternalServerError, message)
}
