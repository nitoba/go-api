package repositories

import "github.com/nitoba/go-api/internal/domain/enterprise/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
