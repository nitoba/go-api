package models

import (
	"time"
)

type UserModel struct {
	ID        string    `gorm:"primaryKey;type:uuid"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"index;unique"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime not null"`
}

func (u *UserModel) TableName() string {
	return "users"
}
