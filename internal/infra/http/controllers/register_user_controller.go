package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/nitoba/go-api/internal/domain/application/use_cases"
	usecases_errors "github.com/nitoba/go-api/internal/domain/application/use_cases/errors"
)

type RegisterUserController struct {
	useCase *usecases.RegisterUseCase
}

func (r *RegisterUserController) Handle(c *gin.Context) {
	var body usecases.RegisterUseCaseRequest
	c.Bind(&body)

	// Validate fields and return errors

	err := r.useCase.Execute(body)

	if errors.Is(err, usecases_errors.ErrUserAlreadyRegistered) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	c.Status(http.StatusCreated)
}

func CreateRegisterUserController(useCase *usecases.RegisterUseCase) *RegisterUserController {
	return &RegisterUserController{
		useCase: useCase,
	}
}
