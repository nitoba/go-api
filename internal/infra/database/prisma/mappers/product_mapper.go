package mappers

import (
	"github.com/nitoba/go-api/internal/domain/enterprise/entity"
	pkg "github.com/nitoba/go-api/pkg/entity"
	"github.com/nitoba/go-api/prisma/db"
)

func ToProductEntity(u *db.ProductModel) *entity.Product {
	id, _ := pkg.ParseID(u.ID)

	return &entity.Product{
		ID:        id,
		Name:      u.Name,
		Price:     u.Price,
		CreatedAt: u.CreatedAt,
	}
}
