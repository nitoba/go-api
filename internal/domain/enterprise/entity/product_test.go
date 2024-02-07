package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct_ShouldBeCreateAnNewProduct(t *testing.T) {
	product, err := NewProduct("Notebook", 2000.0)

	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID.String())
	assert.Equal(t, product.Name, "Notebook")
	assert.Equal(t, product.Price, 2000.0)
	assert.Nil(t, product.Validate())
}

func TestProduct_ShouldNotBeAbleToCreateAnNewProductWithInvalidParams(t *testing.T) {
	product, err := NewProduct("", 2000.0)

	assert.Nil(t, product)
	assert.ErrorIs(t, err, ErrNameIsRequired)

	product, err = NewProduct("Notebook", 0)
	assert.ErrorIs(t, err, ErrPriceIsRequired)

	product, err = NewProduct("Notebook", -1)
	assert.ErrorIs(t, err, ErrInvalidPrice)
}

func TestProduct_ShouldNotBeAbleToCreateANewProductWithPriceEqualsZero(t *testing.T) {
	product, err := NewProduct("Notebook", 0.0)

	assert.Nil(t, product)
	assert.Error(t, err, "price must be positive")
}
