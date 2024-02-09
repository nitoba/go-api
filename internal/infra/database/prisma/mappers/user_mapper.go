package mappers

import (
	"github.com/nitoba/go-api/internal/domain/enterprise/entity"
	pkg "github.com/nitoba/go-api/pkg/entity"
	"github.com/nitoba/go-api/prisma/db"
)

func ToUserEntity(u *db.UserModel) *entity.User {
	id, _ := pkg.ParseID(u.ID)

	return &entity.User{
		ID:       id,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
