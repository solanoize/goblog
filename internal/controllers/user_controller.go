package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/solanoize/goblog/internal/dtos"
	"github.com/solanoize/goblog/internal/usecases"
	"github.com/solanoize/goblog/internal/utils"
)

type UserController interface {
	Register(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	Me(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	UserUseCase usecases.UserUseCase
}

// List implements [UserController].
func (u *userController) List(w http.ResponseWriter, r *http.Request) {
	var err error
	var page int
	var limit int
	var search string
	var count int64
	var paginationResponseDTO dtos.PaginationResponseDTO
	var pagination utils.Pagination = utils.NewPagination(r)

	limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
	page, _ = strconv.Atoi(r.URL.Query().Get("page"))
	search = r.URL.Query().Get("search")

	paginationResponseDTO.Results, paginationResponseDTO.Count, err = u.UserUseCase.All(page, limit, search)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"detail": err.Error()})
		return
	}

	paginationResponseDTO.Next, paginationResponseDTO.Previous = pagination.Paginate(count)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(paginationResponseDTO)

}

// Me implements [UserController].
func (u *userController) Me(w http.ResponseWriter, r *http.Request) {
	var userResponseDTO dtos.UserResponseDTO

	userResponseDTO, _ = r.Context().Value("dto").(dtos.UserResponseDTO)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userResponseDTO)
}

// SignIn implements [UserController].
func (u *userController) SignIn(w http.ResponseWriter, r *http.Request) {
	var err error
	var userSignInDTO dtos.UserSignInDTO
	var userTokenResponseDTO dtos.UserTokenResponseDTO

	err = json.NewDecoder(r.Body).Decode(&userSignInDTO)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"detail": "Invalid json payload"})
		return
	}

	err = validator.New().Struct(userSignInDTO)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.NewValidation().Format(err))
		return
	}

	userTokenResponseDTO, err = u.UserUseCase.SignIn(userSignInDTO.Email, userSignInDTO.Password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"detail": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userTokenResponseDTO)

}

// Register implements [UserController].
func (u *userController) Register(w http.ResponseWriter, r *http.Request) {
	var err error = nil
	var userRegisterDTO dtos.UserRegisterDTO
	var userResponseDTO dtos.UserResponseDTO

	err = json.NewDecoder(r.Body).Decode(&userRegisterDTO)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"detail": "Invalid json payload"})
		return
	}

	err = validator.New().Struct(userRegisterDTO)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.NewValidation().Format(err))
		return
	}

	userResponseDTO, err = u.UserUseCase.Register(userRegisterDTO)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"detail": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(userResponseDTO)
}

func NewUserController(userUseCase usecases.UserUseCase) UserController {
	return &userController{UserUseCase: userUseCase}
}
