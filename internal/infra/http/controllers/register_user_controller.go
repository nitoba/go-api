package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
	usecases "github.com/nitoba/apis/internal/domain/application/use_cases"
	usecases_errors "github.com/nitoba/apis/internal/domain/application/use_cases/errors"
)

type RegisterUserController struct {
	useCase *usecases.RegisterUseCase
}

func (r *RegisterUserController) Handle(c echo.Context) error {
	var body usecases.RegisterUseCaseRequest
	c.Bind(&body)

	// Validate fields and return errors

	err := r.useCase.Execute(body)

	if errors.Is(err, usecases_errors.ErrUserAlreadyRegistered) {
		return c.JSON(http.StatusConflict, map[string]string{"message": err.Error()})
	}

	return c.NoContent(http.StatusCreated)
}

func CreateRegisterUserController(useCase *usecases.RegisterUseCase) *RegisterUserController {
	return &RegisterUserController{
		useCase: useCase,
	}
}
