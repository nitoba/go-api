package prisma

import (
	"github.com/nitoba/go-api/configs"
	"github.com/nitoba/go-api/prisma/db"
)

var database *db.PrismaClient

func setupDB() error {
	logger := configs.GetLogger("prisma")
	logger.Info("connecting with postgres")

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	database = client
	return nil
}

func Connect() error {
	if err := setupDB(); err != nil {
		return err
	}
	return nil
}

func GetDB() *db.PrismaClient {
	return database
}
