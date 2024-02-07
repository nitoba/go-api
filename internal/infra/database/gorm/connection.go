package gorm

import (
	"fmt"

	"github.com/nitoba/apis/configs"
	"github.com/nitoba/apis/internal/infra/database/gorm/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB

func open() error {
	if configs.Config.DBDriver == "sqlite" {
		return connectWithSqlite()
	}

	if configs.Config.DBDriver == "postgres" {
		return connectWithPostgres()
	}
	panic("unsupported database driver")
}

func connectWithSqlite() error {
	db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	DBInstance = db

	return nil
}

func connectWithPostgres() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configs.Config.DBHost, configs.Config.DBPort, configs.Config.DBUser, configs.Config.DBPassword, configs.Config.DBName,
	)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return err
	}
	DBInstance = db

	return nil
}

func configureModels() error {
	return DBInstance.AutoMigrate(&models.UserModel{}, &models.ProductModel{})
}

func NewConnectionDB() error {
	if err := open(); err != nil {
		return err
	}

	if err := configureModels(); err != nil {
		return nil
	}

	return nil
}
