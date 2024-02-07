package repositories

import (
	"testing"

	"github.com/nitoba/go-api/internal/infra/database/gorm/models"
	"github.com/nitoba/go-api/test/factories"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Test_UserRepositoryCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&models.UserModel{})
	userRepo := NewUserRepository(db)

	user := factories.MakeUser()
	err = userRepo.Create(user)

	assert.Nil(t, err)
	var userFound models.UserModel
	err = db.First(&userFound).Error

	assert.Nil(t, err)

	assert.Equal(t, user.ID.String(), userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)
}

func Test_UserRepositoryFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&models.UserModel{})
	userRepo := NewUserRepository(db)

	user := factories.MakeUser(map[string]interface{}{"email": "johndoe@gmail.com"})

	userRepo.Create(user)
	user, err = userRepo.FindByEmail("johndoe@gmail.com")
	assert.Nil(t, err)
	assert.Equal(t, user.Email, "johndoe@gmail.com")
}
