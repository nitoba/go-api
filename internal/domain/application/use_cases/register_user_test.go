package usecases

import (
	"testing"

	usecases_errors "github.com/nitoba/apis/internal/domain/application/use_cases/errors"
	cryptography_test "github.com/nitoba/apis/test/cryptography"
	"github.com/nitoba/apis/test/factories"
	"github.com/nitoba/apis/test/repositories"
	"github.com/stretchr/testify/assert"
)

var userRepository *repositories.InMemoryUserRepository
var hashGenerator *cryptography_test.FakeHasher

func makeSut() *RegisterUseCase {
	userRepository = &repositories.InMemoryUserRepository{}
	hashGenerator = &cryptography_test.FakeHasher{}
	return CreateRegisterUseCase(userRepository, hashGenerator)
}

func TestRegisterUseCase(t *testing.T) {

	t.Run("Should be able to register a new user", func(t *testing.T) {
		sut := makeSut()

		res := sut.Execute(RegisterUseCaseRequest{
			Name:     "John Doe",
			Email:    "john.doe@gmail.com",
			Password: "password",
		})

		assert.Nil(t, res)
		assert.Equal(t, userRepository.Users[0].Password, "hashed:password")
	})

	t.Run("Should not be able to register a new user if already exists", func(t *testing.T) {
		sut := makeSut()

		user := factories.MakeUser(map[string]interface{}{"email": "john.doe@gmail.com"})
		userRepository.Users = append(userRepository.Users, user)

		res := sut.Execute(RegisterUseCaseRequest{
			Name:     "John Doe",
			Email:    "john.doe@gmail.com",
			Password: "password",
		})

		assert.ErrorIs(t, res, usecases_errors.ErrUserAlreadyRegistered)
	})
}
