package utils

import (
	"encoding/json"
	"net/http"
)

type Response interface {
	SetStatus(status int) *response
	SetData(data interface{}) *response
	JSON()
}

type response struct {
	responseWriter http.ResponseWriter
	status         int
	data           interface{}
}

// GetResponse implements [Response].
func (r *response) JSON() {
	r.responseWriter.Header().Set("Content-Type", "application/json")
	r.responseWriter.WriteHeader(r.status)
	json.NewEncoder(r.responseWriter).Encode(r.data)
}

// SetData implements [Response].
func (r *response) SetData(data interface{}) *response {
	r.data = data
	return r
}

// SetStatus implements [Response].
func (r *response) SetStatus(status int) *response {
	r.status = status
	return r
}

func NewResponse(w http.ResponseWriter) Response {
	return &response{responseWriter: w}
}
