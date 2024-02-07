package models

import (
	"time"
)

type ProductModel struct {
	ID        string    `gorm:"primaryKey;type:uuid"`
	Name      string    `gorm:"not null"`
	Price     float64   `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null autoCreateTime"`
}

func (u *ProductModel) TableName() string {
	return "products"
}
