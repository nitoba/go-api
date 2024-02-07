package gorm

import (
	"github.com/nitoba/go-api/configs"
	"gorm.io/gorm"
)

var db *gorm.DB

func setupDB() error {
	config := configs.GetConfig()
	if config.DBDriver == "sqlite" {
		database, err := initializeSQLite()
		if err != nil {
			return err
		}
		db = database
		return nil
	}

	if config.DBDriver == "postgres" {
		return initializePostgres()
	}
	panic("unsupported database driver")
}

func Connect() error {
	if err := setupDB(); err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}
