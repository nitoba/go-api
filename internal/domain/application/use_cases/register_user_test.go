package usecases

import (
	"testing"

	usecases_errors "github.com/nitoba/go-api/internal/domain/application/use_cases/errors"
	cryptography_test "github.com/nitoba/go-api/test/cryptography"
	"github.com/nitoba/go-api/test/factories"
	"github.com/nitoba/go-api/test/repositories"
	"github.com/stretchr/testify/assert"
)

type TestRegisterUseCaseConfig struct {
	sut            *RegisterUseCase
	userRepository *repositories.InMemoryUserRepository
	hashGenerator  *cryptography_test.FakeHasher
}

func makeRegisterUseCase() TestRegisterUseCaseConfig {
	userRepository := &repositories.InMemoryUserRepository{}
	hashGenerator := &cryptography_test.FakeHasher{}
	sut := CreateRegisterUseCase(userRepository, hashGenerator)

	return TestRegisterUseCaseConfig{
		sut:            sut,
		userRepository: userRepository,
		hashGenerator:  hashGenerator,
	}
}

func TestRegisterUseCase(t *testing.T) {
	t.Run("Should be able to register a new user", func(t *testing.T) {
		usecase := makeRegisterUseCase()

		res := usecase.sut.Execute(RegisterUseCaseRequest{
			Name:     "John Doe",
			Email:    "john.doe@gmail.com",
			Password: "password",
		})

		assert.Nil(t, res)
		assert.Equal(t, usecase.userRepository.Users[0].Password, "hashed:password")
	})

	t.Run("Should not be able to register a new user if already exists", func(t *testing.T) {
		usecase := makeRegisterUseCase()

		user := factories.MakeUser(map[string]interface{}{"email": "john.doe@gmail.com"})
		usecase.userRepository.Users = append(usecase.userRepository.Users, user)

		res := usecase.sut.Execute(RegisterUseCaseRequest{
			Name:     "John Doe",
			Email:    "john.doe@gmail.com",
			Password: "password",
		})

		assert.ErrorIs(t, res, usecases_errors.ErrUserAlreadyRegistered)
	})
}
