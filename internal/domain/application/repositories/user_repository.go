package repositories

import "github.com/nitoba/apis/internal/domain/enterprise/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
