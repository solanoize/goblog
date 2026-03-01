package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/solanoize/goblog/internal/dtos"
	"github.com/solanoize/goblog/internal/usecases"
	"github.com/solanoize/goblog/internal/utils"
)

type PostController interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type postController struct {
	PostUseCase usecases.PostUseCase
}

// Create implements [PostController].
func (p *postController) Create(w http.ResponseWriter, r *http.Request) {
	var err error
	var postCreateDTO dtos.PostCreateDTO
	var postResponseDTO dtos.PostResponseDTO

	err = json.NewDecoder(r.Body).Decode(&postCreateDTO)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"detail": "Invalid json payload"})
		return
	}

	err = validator.New().Struct(postCreateDTO)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.NewValidation().Format(err))
		return
	}

	postResponseDTO, err = p.PostUseCase.Create(r.Context(), postCreateDTO)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"detail": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(postResponseDTO)
}

func NewPostController(postUseCase usecases.PostUseCase) PostController {
	return &postController{PostUseCase: postUseCase}
}
