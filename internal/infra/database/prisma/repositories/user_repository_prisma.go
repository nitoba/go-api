package repositories

import (
	"fmt"

	"github.com/nitoba/go-api/configs"
	"github.com/nitoba/go-api/internal/domain/enterprise/entity"
	"github.com/nitoba/go-api/internal/infra/database/prisma/mappers"
	"github.com/nitoba/go-api/prisma/db"
)

type UserRepositoryPrisma struct {
	db *db.PrismaClient
}

func (r *UserRepositoryPrisma) Create(user *entity.User) error {
	var config = configs.GetConfig()

	fmt.Printf("user id: %v", user.ID.String())

	if u, err := r.db.User.CreateOne(
		db.User.Name.Set(user.Name),
		db.User.Email.Set(user.Email),
		db.User.Password.Set(user.Password),
		db.User.ID.Set(user.ID.String()),
	).Exec(config.Ctx); err != nil && u == nil {
		return err
	}

	return nil
}

func (r *UserRepositoryPrisma) FindByEmail(email string) (*entity.User, error) {
	var config = configs.GetConfig()
	u, err := r.db.User.FindUnique(db.User.Email.Equals(email)).Exec(config.Ctx)
	if err != nil {
		return nil, err
	}

	return mappers.ToUserEntity(u), nil
}

func (r *UserRepositoryPrisma) FindByID(id string) (*entity.User, error) {
	var config = configs.GetConfig()
	u, err := r.db.User.FindUnique(db.User.ID.Equals(id)).Exec(config.Ctx)
	if err != nil {
		return nil, err
	}

	return mappers.ToUserEntity(u), nil
}

func NewUserRepositoryPrisma(db *db.PrismaClient) *UserRepositoryPrisma {
	return &UserRepositoryPrisma{db: db}
}
