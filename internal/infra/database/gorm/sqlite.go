package gorm

import (
	"os"

	"github.com/nitoba/go-api/configs"
	"github.com/nitoba/go-api/internal/infra/database/gorm/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initializeSQLite() (*gorm.DB, error) {
	logger := configs.GetLogger("sqlite")
	dbPath := "./internal/infra/database/sqlite"
	dbAbsolutePath := dbPath + "/dev.db"

	println(dbAbsolutePath)

	// Check if the database file exists
	_, err := os.Stat(dbAbsolutePath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		// Create the database file and directory
		err = os.MkdirAll(dbPath, os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbAbsolutePath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbAbsolutePath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return db, err
	}

	// Migrate the Schema
	err = db.AutoMigrate(models.DbModels...)
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return db, err
	}
	// Return the DB
	return db, nil
}
