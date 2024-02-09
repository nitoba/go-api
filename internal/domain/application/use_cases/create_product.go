package usecases

import (
	"github.com/nitoba/go-api/internal/domain/application/repositories"
	"github.com/nitoba/go-api/internal/domain/enterprise/entity"
)

type CreateProductUseCase struct {
	repository repositories.ProductRepository
}

func (c *CreateProductUseCase) Execute(name string, price float64, userId string) error {
	product, err := entity.NewProduct(name, price)

	if err != nil {
		return err
	}

	err = c.repository.Create(product, userId)

	if err != nil {
		return err
	}

	return nil
}

func NewCreateProductUseCase(repository repositories.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		repository: repository,
	}
}
