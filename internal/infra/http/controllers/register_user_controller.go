package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/nitoba/go-api/internal/domain/application/use_cases"
	usecases_errors "github.com/nitoba/go-api/internal/domain/application/use_cases/errors"
	"github.com/nitoba/go-api/internal/infra/http/validations"
)

type RegisterUserController struct {
	useCase *usecases.RegisterUseCase
}

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required|min_len:6" message:"required:{field} is required" label:"User Name"`
	Email    string `json:"email" validate:"email" message:"email is invalid" label:"User Email"`
	Password string `json:"password" validate:"required|min_len:6" message:"required:{field} is required" label:"Password"`
}

func (r *RegisterUserController) Handle(c *gin.Context) {
	var body RegisterUserRequest
	c.Bind(&body)

	if !validations.SendBadRequestValidation(body, c) {
		return
	}

	err := r.useCase.Execute(usecases.RegisterUseCaseRequest{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	})

	if errors.Is(err, usecases_errors.ErrUserAlreadyRegistered) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func CreateRegisterUserController(useCase *usecases.RegisterUseCase) *RegisterUserController {
	return &RegisterUserController{
		useCase: useCase,
	}
}
