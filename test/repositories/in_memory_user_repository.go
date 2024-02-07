package repositories

import "github.com/nitoba/go-api/internal/domain/enterprise/entity"

type InMemoryUserRepository struct {
	Users []*entity.User
}

func (r *InMemoryUserRepository) Create(user *entity.User) error {
	r.Users = append(r.Users, user)
	return nil
}

func (r *InMemoryUserRepository) FindByEmail(email string) (*entity.User, error) {
	for _, user := range r.Users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, nil
}

func (r *InMemoryUserRepository) FindByID(id string) (*entity.User, error) {
	for _, user := range r.Users {
		if user.ID.String() == id {
			return user, nil
		}
	}

	return nil, nil
}
