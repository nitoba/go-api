package repositories

import "github.com/nitoba/go-api/internal/domain/enterprise/entity"

type ProductRepository interface {
	Create(product *entity.Product) error
	FindByID(id string) (*entity.Product, error)
	Find(page, limit int, sort string) ([]*entity.Product, error)
	Delete(id string) error
	Update(product *entity.Product) error
}
