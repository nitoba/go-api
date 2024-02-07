package repositories

import (
	"testing"

	"github.com/nitoba/apis/internal/domain/enterprise/entity"
	"github.com/nitoba/apis/internal/infra/database/gorm/models"
	"github.com/nitoba/apis/test/factories"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Test_ProductRepositoryCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&models.ProductModel{})
	productRepo := NewProductRepository(db)

	product := factories.MakeProduct()

	err = productRepo.Create(product)

	assert.Nil(t, err)
	var productFound models.ProductModel
	err = db.First(&productFound).Error

	assert.Nil(t, err)

	assert.Equal(t, product.ID.String(), productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func Test_ProductRepositoryFindProductById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&models.ProductModel{})
	productRepo := NewProductRepository(db)

	product := factories.MakeProduct()

	err = productRepo.Create(product)
	assert.Nil(t, err)

	productFound, err := productRepo.FindById(product.ID.String())

	assert.Nil(t, err)

	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func Test_ProductRepositoryShouldReturnAErrorIfNotFoundAProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&models.ProductModel{})
	productRepo := NewProductRepository(db)

	product := factories.MakeProduct()

	_, err = productRepo.FindById(product.ID.String())

	assert.Error(t, err)
}

func Test_ProductRepositoryFindProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&models.ProductModel{})
	productRepo := NewProductRepository(db)

	products := []*entity.Product{factories.MakeProduct(), factories.MakeProduct(), factories.MakeProduct()}

	for _, product := range products {
		err = productRepo.Create(product)
		assert.Nil(t, err)
	}

	productsFound, err := productRepo.Find(1, 10, "ASC")
	assert.Nil(t, err)
	assert.Len(t, productsFound, 3)
}

func Test_ProductRepositoryFindPaginatedProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&models.ProductModel{})
	productRepo := NewProductRepository(db)

	for i := 0; i < 14; i++ {
		product := factories.MakeProduct()
		err = productRepo.Create(product)
	}

	productsFound, err := productRepo.Find(1, 10, "DESC")
	assert.Nil(t, err)
	assert.Len(t, productsFound, 10)

	productsFound, err = productRepo.Find(2, 10, "DESC")
	assert.Nil(t, err)
	assert.Len(t, productsFound, 4)
	assert.True(t, productsFound[0].CreatedAt.After(productsFound[1].CreatedAt))
	assert.True(t, productsFound[1].CreatedAt.After(productsFound[2].CreatedAt))
	assert.True(t, productsFound[2].CreatedAt.After(productsFound[3].CreatedAt))
}

func Test_ProductRepositoryShouldReturnAnEmptyListIfNotExistsProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&models.ProductModel{})
	productRepo := NewProductRepository(db)

	productsFound, err := productRepo.Find(1, 10, "ASC")
	assert.Nil(t, err)
	assert.Len(t, productsFound, 0)
}

func Test_ProductRepositoryShouldDeleteAProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&models.ProductModel{})
	productRepo := NewProductRepository(db)

	product := factories.MakeProduct()

	err = productRepo.Create(product)

	assert.Nil(t, err)

	err = productRepo.Delete(product.ID.String())

	assert.Nil(t, err)

	err = db.First(&product).Error

	assert.Error(t, err)
}

func Test_ProductRepositoryUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&models.ProductModel{})
	productRepo := NewProductRepository(db)

	product := factories.MakeProduct()

	err = productRepo.Create(product)

	assert.Nil(t, err)

	product.Name = "New Product Name"
	product.Price = 400.0

	err = productRepo.Update(product)
	assert.Nil(t, err)

	var productFound models.ProductModel

	err = db.Where("id = ?", product.ID.String()).First(&productFound).Error

	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)

}
