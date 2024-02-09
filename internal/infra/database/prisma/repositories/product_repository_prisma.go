package repositories

import (
	"github.com/nitoba/go-api/configs"
	"github.com/nitoba/go-api/internal/domain/enterprise/entity"
	"github.com/nitoba/go-api/internal/infra/database/prisma/mappers"
	"github.com/nitoba/go-api/prisma/db"
)

type ProductRepositoryPrisma struct {
	db *db.PrismaClient
}

func (r *ProductRepositoryPrisma) Create(product *entity.Product, userID string) error {
	var config = configs.GetConfig()
	_, err := r.db.Product.CreateOne(
		db.Product.Name.Set(product.Name),
		db.Product.Price.Set(product.Price),
		db.Product.User.Link(db.User.ID.Equals(userID)),
		db.Product.ID.Set(product.ID.String()),
	).Exec(config.Ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryPrisma) Update(product *entity.Product) error {
	var config = configs.GetConfig()
	_, err := r.db.Product.FindUnique(db.Product.ID.Equals(product.ID.String())).Update(
		db.Product.Name.Set(product.Name),
		db.Product.Price.Set(product.Price),
	).Exec(config.Ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepositoryPrisma) FindByID(id string) (*entity.Product, error) {
	var config = configs.GetConfig()
	u, err := r.db.Product.FindUnique(db.Product.ID.Equals(id)).Exec(config.Ctx)
	if err != nil {
		return nil, err
	}
	return mappers.ToProductEntity(u), nil
}

func (r *ProductRepositoryPrisma) Find(page, limit int, sort string) ([]*entity.Product, error) {
	var config = configs.GetConfig()
	var products []*entity.Product = []*entity.Product{}
	p, err := r.db.Product.FindMany().Exec(config.Ctx)
	if err != nil {
		return nil, err
	}

	for _, product := range p {
		products = append(products, mappers.ToProductEntity(&product))
	}

	return products, nil
}

func (r *ProductRepositoryPrisma) Delete(id string) error {
	var config = configs.GetConfig()
	_, err := r.db.Product.FindUnique(db.Product.ID.Equals(id)).Delete().Exec(config.Ctx)
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepositoryPrisma(db *db.PrismaClient) *ProductRepositoryPrisma {
	return &ProductRepositoryPrisma{db: db}
}
