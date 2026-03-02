package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/solanoize/goblog/internal/utils"
)

type Controller interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Retrieve(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	Service Service
}

// SignIn implements [Controller].
func (c *controller) SignIn(w http.ResponseWriter, r *http.Request) {
	requestContract := SignInRequestContract{}

	err := json.NewDecoder(r.Body).Decode(&requestContract)
	if err != nil {
		utils.RenderBadRequest(w, "Invalid json payload")
		return
	}

	err = validator.New().Struct(requestContract)
	if err != nil {
		utils.RenderBadRequest(w, err.Error())
		return
	}

	responseContract, err := c.Service.SignIn(r.Context(), requestContract)
	if err != nil {
		utils.RenderBadRequest(w, err.Error())
		return
	}

	utils.RenderOK(w, responseContract)
}

// Create implements [Controller].
func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	requestContract := SignUpRequestContract{}

	err := json.NewDecoder(r.Body).Decode(&requestContract)
	if err != nil {
		utils.RenderBadRequest(w, "Invalid json payload")
		return
	}

	err = validator.New().Struct(requestContract)
	if err != nil {
		utils.RenderBadRequest(w, err.Error())
		return
	}

	responseContract, err := c.Service.SignUp(r.Context(), requestContract)
	if err != nil {
		utils.RenderForbidden(w, err.Error())
		return
	}

	utils.RenderCreated(w, responseContract)
}

// Destroy implements [Controller].
func (c *controller) Destroy(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// List implements [Controller].
func (c *controller) List(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Retrieve implements [Controller].
func (c *controller) Retrieve(w http.ResponseWriter, r *http.Request) {
	claimContract, err := c.Service.Auth().GetClaim(r)
	if err != nil {
		utils.RenderUnauthorized(w, err.Error())
		return
	}

	responseContract, err := c.Service.GetByID(r.Context(), claimContract.UserID)
	if err != nil {
		utils.RenderNotFound(w, err.Error())
		return
	}

	utils.RenderOK(w, responseContract)
}

// Update implements [Controller].
func (c *controller) Update(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func NewController(s Service) Controller {
	return &controller{Service: s}
}
