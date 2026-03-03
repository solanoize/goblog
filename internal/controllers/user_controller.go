package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/solanoize/goblog/internal/dto"
	"github.com/solanoize/goblog/internal/globals"
	"github.com/solanoize/goblog/internal/models"
	"github.com/solanoize/goblog/internal/services"
	"github.com/solanoize/goblog/internal/utils"
)

type UserController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	UserService services.UserService
}

// SignUp implements [UserController].
func (u *userController) SignUp(w http.ResponseWriter, r *http.Request) {
	requestDTO := dto.SignUpRequestDTO{}

	fmt.Println("Go DB = ", globals.GlobalDB)
	if globals.GlobalDB != nil {
		fmt.Println("Database connection no nil!")
		err := globals.GlobalDB.Create(&models.User{
			Username: "uhuy",
			Email:    "uhuy@mail.com",
			Password: "123qwe",
		}).Error

		if err != nil {
			fmt.Println("Error create user")
		}
	}

	err := json.NewDecoder(r.Body).Decode(&requestDTO)
	if err != nil {
		utils.BadRequestRender(w, err.Error())
		return
	}

	err = validator.New().Struct(requestDTO)
	if err != nil {
		utils.BadRequestRender(w, err.Error())
		return
	}

	responseDTO, err := u.UserService.SignUp(r.Context(), requestDTO)
	if err != nil {
		utils.ForbiddenRender(w, err.Error())
		return
	}

	utils.OKRender(w, responseDTO)
}

func NewUserController(userService services.UserService) UserController {
	return &userController{UserService: userService}
}
