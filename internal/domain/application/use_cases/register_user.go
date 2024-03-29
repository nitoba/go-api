package usecases

import (
	"github.com/nitoba/go-api/internal/domain/application/cryptography"
	"github.com/nitoba/go-api/internal/domain/application/repositories"
	usecases_errors "github.com/nitoba/go-api/internal/domain/application/use_cases/errors"
	"github.com/nitoba/go-api/internal/domain/enterprise/entity"
)

type RegisterUseCaseRequest struct {
	Name     string
	Email    string
	Password string
}

type RegisterUseCaseResponse error

type RegisterUseCase struct {
	cryptography   cryptography.HashGenerator
	userRepository repositories.UserRepository
}

func (u *RegisterUseCase) Execute(request RegisterUseCaseRequest) RegisterUseCaseResponse {
	userExists, _ := u.userRepository.FindByEmail(request.Email)

	if userExists != nil {
		return usecases_errors.ErrUserAlreadyRegistered
	}

	hashedPassword, err := u.cryptography.Hash(request.Password)

	if err != nil {
		return err
	}

	user, err := entity.NewUser(request.Name, request.Email, hashedPassword)

	if err != nil {
		return err
	}

	return u.userRepository.Create(user)
}

func CreateRegisterUseCase(userRepository repositories.UserRepository, cryptography cryptography.HashGenerator) *RegisterUseCase {
	return &RegisterUseCase{
		userRepository: userRepository,
		cryptography:   cryptography,
	}
}
