package repositories

import (
	"errors"

	"github.com/nitoba/go-api/internal/domain/enterprise/entity"
	"github.com/nitoba/go-api/internal/infra/database/gorm/mappers"
	"github.com/nitoba/go-api/internal/infra/database/gorm/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepositoryGorm struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepositoryGorm {
	return ProductRepositoryGorm{DB: db}
}

func (r *ProductRepositoryGorm) Create(product *entity.Product) error {
	productModel := mappers.ToProductDBModel(product)
	return r.DB.Create(&productModel).Error
}

func (r *ProductRepositoryGorm) FindById(id string) (*entity.Product, error) {
	var productModel models.ProductModel

	err := r.DB.Model(&productModel).Where("id = ?", id).First(&productModel).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Product Not Found")
	}
	return mappers.ToProductEntity(&productModel), nil
}

func (r *ProductRepositoryGorm) Update(product *entity.Product) error {
	productModel := mappers.ToProductDBModel(product)

	err := r.DB.Where("id = ?", product.ID.String()).First(&models.ProductModel{}).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		println("Product Not Found")

		return err
	}

	return r.DB.Save(&productModel).Error
}

func (r *ProductRepositoryGorm) Find(page, limit int, sort string) ([]*entity.Product, error) {
	var products []*models.ProductModel

	offset := (page - 1) * limit

	isDesc := true

	if sort == "ASC" {
		isDesc = false
	}

	orderBy := clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: isDesc}

	err := r.DB.Model(&products).Offset(offset).Limit(limit).Order(orderBy).Find(&products).Error

	if err != nil {
		return nil, err
	}

	var result []*entity.Product = []*entity.Product{}

	for _, product := range products {
		result = append(result, mappers.ToProductEntity(product))
	}

	return result, nil
}

func (r *ProductRepositoryGorm) Delete(id string) error {
	var productModel models.ProductModel
	return r.DB.Where("id = ?", id).Delete(&productModel).Error
}
