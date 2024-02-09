package repositories

import (
	"context"
	"fmt"

	"github.com/nitoba/go-api/internal/domain/enterprise/entity"
	"github.com/nitoba/go-api/internal/infra/database/prisma/mappers"
	"github.com/nitoba/go-api/prisma/db"
)

type UserRepositoryPrisma struct {
	db *db.PrismaClient
}

var ctx = context.Background()

func (r *UserRepositoryPrisma) Create(user *entity.User) error {
	if u, err := r.db.User.CreateOne(
		db.User.Name.Set(user.Name),
		db.User.Email.Set(user.Email),
		db.User.Password.Set(user.Password),
	).Exec(ctx); err != nil && u == nil {
		return err
	} else {
		fmt.Printf("User created: %v\n", user)
	}

	return nil
}

func (r *UserRepositoryPrisma) FindByEmail(email string) (*entity.User, error) {
	u, err := r.db.User.FindUnique(db.User.Email.Equals(email)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return mappers.ToUserEntity(u), nil
}

func (r *UserRepositoryPrisma) FindByID(id string) (*entity.User, error) {
	u, err := r.db.User.FindUnique(db.User.ID.Equals(id)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return mappers.ToUserEntity(u), nil
}

func NewUserRepositoryPrisma(db *db.PrismaClient) *UserRepositoryPrisma {
	return &UserRepositoryPrisma{db: db}
}
