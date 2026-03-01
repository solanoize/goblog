package usecases

import (
	"context"
	"errors"
	"log"

	"github.com/solanoize/goblog/internal/dtos"
	"github.com/solanoize/goblog/internal/mappers"
	"github.com/solanoize/goblog/internal/models"
	"github.com/solanoize/goblog/internal/recontext"
	"github.com/solanoize/goblog/internal/repositories"
)

type PostUseCase interface {
	Create(ctx context.Context, postCreateDTO dtos.PostCreateDTO) (dtos.PostResponseDTO, error)
}

type postUseCase struct {
	Logger         *log.Logger
	PostRepository repositories.PostRepository
}

// Create implements [PostUseCase].
func (p *postUseCase) Create(ctx context.Context, postCreateDTO dtos.PostCreateDTO) (dtos.PostResponseDTO, error) {
	var err error
	var post models.Post
	var userContextDTO dtos.UserContextDTO
	var postResponseDTO dtos.PostResponseDTO

	userContextDTO = recontext.NewUserContext().GetUserContext(ctx)

	p.Logger.Println(userContextDTO)
	post = mappers.NewPostCreateMapper().ToModel(postCreateDTO)
	post.UserID = userContextDTO.ID
	post, err = p.PostRepository.Create(ctx, post)
	postResponseDTO = mappers.NewPostResponseMapper().ToDTO(post)
	if err != nil {
		p.Logger.Println(err)
		return postResponseDTO, errors.New("Post tidak berhasil ditambahkan.")
	}

	return postResponseDTO, nil
}

func NewPostUseCase(logger *log.Logger, postRespository repositories.PostRepository) PostUseCase {
	return &postUseCase{
		Logger:         logger,
		PostRepository: postRespository,
	}
}
