package usecases

import (
	"github.com/nitoba/go-api/internal/domain/application/cryptography"
	"github.com/nitoba/go-api/internal/domain/application/repositories"
	usecases_errors "github.com/nitoba/go-api/internal/domain/application/use_cases/errors"
)

type AuthenticateUseCase struct {
	userRepository repositories.UserRepository
	cryptography   cryptography.HashComparer
}

type AuthenticateUseCaseRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateUseCaseResponse struct {
	Token string `json:"token"`
}

func (c *AuthenticateUseCase) Execute(request AuthenticateUseCaseRequest) (*AuthenticateUseCaseResponse, error) {
	user, err := c.userRepository.FindByEmail(request.Email)

	if err != nil || user == nil {
		return nil, usecases_errors.ErrWrongCredentials
	}

	passwordsMatch := c.cryptography.Compare(request.Password, user.Password)

	if !passwordsMatch {
		return nil, usecases_errors.ErrWrongCredentials
	}

	// TODO: Generate a token JWT and return it

	return &AuthenticateUseCaseResponse{
		Token: "token",
	}, nil
}

func NewAuthenticate(userRepository repositories.UserRepository, cryptography cryptography.HashComparer) *AuthenticateUseCase {
	return &AuthenticateUseCase{
		userRepository: userRepository,
		cryptography:   cryptography,
	}
}
