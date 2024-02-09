package main

import (
	"fmt"

	"github.com/nitoba/go-api/configs"
	"github.com/nitoba/go-api/internal/infra/database/prisma"
	"github.com/nitoba/go-api/internal/infra/http/server"
)

func main() {
	logger := configs.GetLogger("main")
	if _, err := configs.LoadConfig(); err != nil {
		logger.Errorf("error loading config: %v", err)
		panic(err)
	}

	if err := prisma.Connect(); err != nil {
		logger.Errorf("error to connect with database: %v", err)
		panic(err)
	}

	config := configs.GetConfig()
	server := server.Setup()

	server.Run(fmt.Sprintf(":%s", config.WebServerPort))
}
