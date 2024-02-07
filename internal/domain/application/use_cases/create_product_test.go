package usecases_test

import (
	"testing"

	usecases "github.com/nitoba/go-api/internal/domain/application/use_cases"
	cryptography_test "github.com/nitoba/go-api/test/cryptography"
	"github.com/nitoba/go-api/test/repositories"
	"github.com/stretchr/testify/assert"
)

type TestCreateProductUseCaseConfig struct {
	sut               *usecases.CreateProductUseCase
	productRepository *repositories.InMemoryProductRepository
	hashGenerator     *cryptography_test.FakeHasher
}

func makeCreateProductUseCase() TestCreateProductUseCaseConfig {
	productRepository := &repositories.InMemoryProductRepository{}
	sut := usecases.NewCreateProductUseCase(productRepository)

	return TestCreateProductUseCaseConfig{
		sut:               sut,
		productRepository: productRepository,
	}
}

func TestCreateProductUseCase(t *testing.T) {
	t.Run("Should be able to create a new product", func(t *testing.T) {
		usecase := makeCreateProductUseCase()

		res := usecase.sut.Execute("New Product", 99.0)

		assert.Nil(t, res)
		assert.Equal(t, usecase.productRepository.Products[0].Name, "New Product")
		assert.Equal(t, usecase.productRepository.Products[0].Price, 99.0)
	})

	t.Run("Should not be able to create a new product with incorrect infos", func(t *testing.T) {
		usecase := makeCreateProductUseCase()

		res := usecase.sut.Execute("", 99.0)

		assert.NotNil(t, res)
		assert.Empty(t, usecase.productRepository.Products)

		res = usecase.sut.Execute("Product", 0)

		assert.NotNil(t, res)
		assert.Empty(t, usecase.productRepository.Products)

		res = usecase.sut.Execute("", 0)

		assert.NotNil(t, res)
		assert.Empty(t, usecase.productRepository.Products)
	})
}
