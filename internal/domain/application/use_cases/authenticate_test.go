package usecases

import (
	"testing"

	usecases_errors "github.com/nitoba/go-api/internal/domain/application/use_cases/errors"
	cryptography_test "github.com/nitoba/go-api/test/cryptography"
	"github.com/nitoba/go-api/test/factories"
	"github.com/nitoba/go-api/test/repositories"
	"github.com/stretchr/testify/assert"
)

type TestAuthenticateUseCaseConfig struct {
	sut            *AuthenticateUseCase
	userRepository *repositories.InMemoryUserRepository
	fakeHasher     *cryptography_test.FakeHasher
	fakeEncrypter  *cryptography_test.FakeEncrypter
}

func makeAuthenticateUseCase() TestAuthenticateUseCaseConfig {
	userRepository := &repositories.InMemoryUserRepository{}
	hashComparer := &cryptography_test.FakeHasher{}
	fakeEncrypter := &cryptography_test.FakeEncrypter{}
	sut := NewAuthenticate(userRepository, hashComparer, fakeEncrypter)

	return TestAuthenticateUseCaseConfig{
		sut:            sut,
		userRepository: userRepository,
		fakeHasher:     hashComparer,
		fakeEncrypter:  fakeEncrypter,
	}
}

func TestAuthenticateUseCase(t *testing.T) {
	t.Run("Should be able to authenticate an user", func(t *testing.T) {
		usecase := makeAuthenticateUseCase()

		passwordHashed, _ := usecase.fakeHasher.Hash("password")

		user := factories.MakeUser(map[string]interface{}{"email": "test@example.com", "password": passwordHashed})

		usecase.userRepository.Users = append(usecase.userRepository.Users, user)

		res, err := usecase.sut.Execute(AuthenticateUseCaseRequest{
			Email:    "test@example.com",
			Password: "password",
		})

		println("Token: ", res.Token)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.NotEmpty(t, res.Token)
	})

	t.Run("Should not be able to authenticate an user with wrong credentials", func(t *testing.T) {
		usecase := makeAuthenticateUseCase()

		user := factories.MakeUser()

		usecase.userRepository.Users = append(usecase.userRepository.Users, user)

		res, err := usecase.sut.Execute(AuthenticateUseCaseRequest{
			Email:    "john.doe@gmail.com",
			Password: "password",
		})

		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, usecases_errors.ErrWrongCredentials)
	})

	t.Run("Should not be able to authenticate an user if not exists", func(t *testing.T) {
		usecase := makeAuthenticateUseCase()

		res, err := usecase.sut.Execute(AuthenticateUseCaseRequest{
			Email:    "john.doe@gmail.com",
			Password: "password",
		})

		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, usecases_errors.ErrWrongCredentials)
	})
}
