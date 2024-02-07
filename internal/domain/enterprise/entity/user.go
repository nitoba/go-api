package entity

import (
	"github.com/nitoba/go-api/pkg/entity"
)

type User struct {
	ID       entity.ID
	Name     string
	Email    string
	Password string
}

func NewUser(name, email, password string) (*User, error) {
	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}
