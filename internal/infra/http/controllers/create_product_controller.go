package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/nitoba/go-api/internal/domain/application/use_cases"
	"github.com/nitoba/go-api/internal/infra/http/validations"
)

type CreateProductController struct {
	useCase *usecases.CreateProductUseCase
}

type CreateProductRequest struct {
	Name  string  `json:"name" validate:"required|min_len:6" message:"required:{field} is required"`
	Price float64 `json:"price" validate:"required|float|min:1" message:"float:price must float|min:price min value is 1"`
}

func (r *CreateProductController) Handle(c *gin.Context) {
	var body CreateProductRequest
	c.Bind(&body)

	if !validations.SendBadRequestValidation(body, c) {
		return
	}

	err := r.useCase.Execute(body.Name, body.Price)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func NewCreateProductController(useCase *usecases.CreateProductUseCase) *CreateProductController {
	return &CreateProductController{
		useCase: useCase,
	}
}
