package dtos

type PostCreateDTO struct {
	Title     string `json:"title" validate:"required,min=5"`
	Content   string `json:"content" validate:"required,min=20"`
	Publsihed bool   `json:"published"`
}
