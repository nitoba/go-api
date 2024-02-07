package repositories

import (
	"errors"

	"github.com/nitoba/go-api/internal/domain/enterprise/entity"
	"github.com/nitoba/go-api/internal/infra/database/gorm/mappers"
	"github.com/nitoba/go-api/internal/infra/database/gorm/models"
	"gorm.io/gorm"
)

type UserRepositoryGorm struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryGorm {
	return &UserRepositoryGorm{DB: db}
}

func (u *UserRepositoryGorm) Create(user *entity.User) error {
	userModel := mappers.ToUserDBModel(user)
	return u.DB.Create(&userModel).Error
}

func (u *UserRepositoryGorm) FindByEmail(email string) (*entity.User, error) {
	var user models.UserModel
	err := u.DB.Where("email =?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return mappers.ToUserEntity(&user), nil
}

func (u *UserRepositoryGorm) FindByID(id string) (*entity.User, error) {
	var user models.UserModel
	err := u.DB.Where("id =?", id).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return mappers.ToUserEntity(&user), nil
}
