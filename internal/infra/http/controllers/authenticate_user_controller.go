package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/nitoba/go-api/internal/domain/application/use_cases"
	"github.com/nitoba/go-api/internal/infra/http/validations"
)

type AuthenticateUserController struct {
	useCase *usecases.AuthenticateUseCase
}

type AuthenticateUserRequest struct {
	Email    string `json:"email" validate:"email" message:"email is invalid" label:"User Email"`
	Password string `json:"password" validate:"required|min_len:6" message:"required:{field} is required" label:"Password"`
}

func (r *AuthenticateUserController) Handle(c *gin.Context) {
	var body AuthenticateUserRequest
	c.Bind(&body)

	if !validations.SendBadRequestValidation(body, c) {
		return
	}

	resp, err := r.useCase.Execute(usecases.AuthenticateUseCaseRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": resp.Token,
	})
}

func NewAuthenticateUserController(useCase *usecases.AuthenticateUseCase) *AuthenticateUserController {
	return &AuthenticateUserController{
		useCase: useCase,
	}
}
