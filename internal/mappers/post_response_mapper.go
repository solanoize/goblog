package mappers

import (
	"github.com/solanoize/goblog/internal/dtos"
	"github.com/solanoize/goblog/internal/models"
)

type PostResponseMapper interface {
	ToDTO(post models.Post) dtos.PostResponseDTO
	ToDTOS(posts []models.Post) []dtos.PostResponseDTO
}

type postResponseMapper struct{}

// ToDTOS implements [PostResponseMapper].
func (p *postResponseMapper) ToDTOS(posts []models.Post) []dtos.PostResponseDTO {
	var postResponseDTOS []dtos.PostResponseDTO = make([]dtos.PostResponseDTO, 0, len(posts))
	var post models.Post

	for _, post = range posts {
		postResponseDTOS = append(postResponseDTOS, p.ToDTO(post))
	}

	return postResponseDTOS
}

func (p *postResponseMapper) ToDTO(post models.Post) dtos.PostResponseDTO {

	var postResponseDTO dtos.PostResponseDTO = dtos.PostResponseDTO{
		ID:        post.ID,
		Title:     post.Title,
		Slug:      post.Slug,
		Content:   post.Content,
		Published: post.Published,
		CreatedAt: post.CreatedAt,
	}

	if post.User != nil {
		var userResponseDTO dtos.UserResponseDTO = NewUserResponseMapper().ToDTO(*post.User)
		postResponseDTO.Author = &userResponseDTO
	}

	return postResponseDTO
}

func NewPostResponseMapper() PostResponseMapper {
	return &postResponseMapper{}
}
