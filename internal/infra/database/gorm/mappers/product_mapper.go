package mappers

import (
	"github.com/nitoba/go-api/internal/domain/enterprise/entity"
	"github.com/nitoba/go-api/internal/infra/database/gorm/models"
	pkg "github.com/nitoba/go-api/pkg/entity"
)

func ToProductEntity(u *models.ProductModel) *entity.Product {
	id, _ := pkg.ParseID(u.ID)

	return &entity.Product{
		ID:        id,
		Name:      u.Name,
		Price:     u.Price,
		CreatedAt: u.CreatedAt,
	}
}

func ToProductDBModel(p *entity.Product) *models.ProductModel {
	return &models.ProductModel{
		ID:        p.ID.String(),
		Name:      p.Name,
		Price:     p.Price,
		CreatedAt: p.CreatedAt,
	}
}
