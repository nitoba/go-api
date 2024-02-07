package mappers

import (
	"time"

	"github.com/nitoba/apis/internal/domain/enterprise/entity"
	"github.com/nitoba/apis/internal/infra/database/gorm/models"
	pkg "github.com/nitoba/apis/pkg/entity"
)

func ToUserEntity(u *models.UserModel) *entity.User {
	id, _ := pkg.ParseID(u.ID)

	return &entity.User{
		ID:       id,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}

func ToUserDBModel(u *entity.User) *models.UserModel {
	return &models.UserModel{
		ID:        u.ID.String(),
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: time.Now(),
	}
}