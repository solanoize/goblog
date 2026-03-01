package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Pagination interface {
	Paginate(count int64) (*string, *string)
}

type pagination struct {
	request *http.Request
	page    int
	limit   int
	count   int64
}

// Links implements [Pagination].
func (p *pagination) Paginate(count int64) (*string, *string) {
	var baseURL string = "http://" + p.request.Host + p.request.URL.Path
	var queryParams url.Values = p.request.URL.Query()

	var next, previous *string

	if int64(p.page*p.limit) < count {
		queryParams.Set("page", fmt.Sprintf("%d", p.page+1))
		queryParams.Set("limit", fmt.Sprintf("%d", p.limit))
		nextURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())
		next = &nextURL
	}

	if p.page > 1 {
		queryParams.Set("page", fmt.Sprintf("%d", p.page-1))
		queryParams.Set("limit", fmt.Sprintf("%d", p.limit))
		prevURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())
		previous = &prevURL
	}

	return next, previous

}

func NewPagination(r *http.Request) Pagination {
	var limit int
	var page int
	limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
	page, _ = strconv.Atoi(r.URL.Query().Get("page"))

	return &pagination{
		request: r,
		limit:   limit,
		page:    page,
	}
}
