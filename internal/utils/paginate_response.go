package utils

import (
	"encoding/json"
	"net/http"
)

type PaginateResponse interface {
	SetStatus(status int) *paginateResponse
	SetCount(count int64) *paginateResponse
	SetData(data interface{}) *paginateResponse
	SetNextPrevious(next *string, previous *string) *paginateResponse
	JSON()
}

type paginateResponse struct {
	status         int
	count          int64
	next           *string
	previous       *string
	data           interface{}
	responseWriter http.ResponseWriter
}

// SetNextPrevious implements [PaginateResponse].
func (p *paginateResponse) SetNextPrevious(next *string, previous *string) *paginateResponse {
	p.next = next
	p.previous = previous
	return p
}

// JSON implements [PaginateResponse].
func (p *paginateResponse) JSON() {
	p.responseWriter.Header().Set("Content-Type", "application/json")
	p.responseWriter.WriteHeader(p.status)
	json.NewEncoder(p.responseWriter).Encode(map[string]any{
		"next":     p.next,
		"previous": p.previous,
		"count":    p.count,
		"results":  p.data,
	})
}

// SetCount implements [PaginateResponse].
func (p *paginateResponse) SetCount(count int64) *paginateResponse {
	p.count = count
	return p
}

// SetData implements [PaginateResponse].
func (p *paginateResponse) SetData(data interface{}) *paginateResponse {
	p.data = data
	return p
}

// SetStatus implements [PaginateResponse].
func (p *paginateResponse) SetStatus(status int) *paginateResponse {
	p.status = status
	return p
}

func NewPaginateResponse(w http.ResponseWriter) PaginateResponse {
	return &paginateResponse{responseWriter: w}
}
