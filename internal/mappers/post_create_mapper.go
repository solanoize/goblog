package mappers

import (
	"github.com/solanoize/goblog/internal/dtos"
	"github.com/solanoize/goblog/internal/models"
)

type PostCreateMapper interface {
	ToModel(postCreateDTO dtos.PostCreateDTO) models.Post
}

type postCreateMapper struct{}

// ToModel implements [PostCreateMapper].
func (p *postCreateMapper) ToModel(postCreateDTO dtos.PostCreateDTO) models.Post {
	return models.Post{
		Title:     postCreateDTO.Title,
		Content:   postCreateDTO.Content,
		Published: postCreateDTO.Publsihed,
	}
}

func NewPostCreateMapper() PostCreateMapper {
	return &postCreateMapper{}
}
