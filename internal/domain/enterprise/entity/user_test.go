package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_ShouldBeCreateAnNewUser(t *testing.T) {
	user, err := NewUser("Jon Doe", "johndoe@example.com", "password")

	assert.NoError(t, err)
	assert.NotEmpty(t, user.ID.String())
	assert.Equal(t, user.Name, "Jon Doe")
	assert.Equal(t, user.Email, "johndoe@example.com")
	assert.Equal(t, user.Password, "password")
}
