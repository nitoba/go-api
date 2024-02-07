package gorm

import (
	"fmt"

	"github.com/nitoba/go-api/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializePostgres() error {
	config := configs.GetConfig()
	logger := configs.GetLogger("postgres")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName,
	)

	database, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error to initializing postgres: %v", err)
		return err
	}
	db = database

	return nil
}
