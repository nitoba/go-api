package repositories

import "github.com/nitoba/go-api/internal/domain/enterprise/entity"

type InMemoryProductRepository struct {
	Products []*entity.Product
}

func (r *InMemoryProductRepository) Create(product *entity.Product) error {
	r.Products = append(r.Products, product)
	return nil
}

func (r *InMemoryProductRepository) FindByID(id string) (*entity.Product, error) {
	return nil, nil
}
func (r *InMemoryProductRepository) Find(page, limit int, sort string) (*[]entity.Product, error) {
	return nil, nil
}
func (r *InMemoryProductRepository) Delete(id string) error {
	return nil
}
func (r *InMemoryProductRepository) Update(product *entity.Product) error {
	return nil
}
