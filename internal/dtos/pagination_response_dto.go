package dtos

type PaginationResponseDTO struct {
	Count    int64       `json:"count"`
	Previous *string     `json:"previous"`
	Next     *string     `json:"next"`
	Results  interface{} `json:"results"`
}
