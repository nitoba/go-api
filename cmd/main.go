package main

import (
	"fmt"

	"github.com/nitoba/go-api/configs"
	"github.com/nitoba/go-api/internal/infra/database/prisma"
	"github.com/nitoba/go-api/internal/infra/http/server"
)

func main() {
	logger := configs.GetLogger("main")
	conf, err := configs.LoadConfig()

	if err != nil {
		logger.Errorf("error loading config: %v", err)
		panic(err)
	}

	err = prisma.Connect()

	if err != nil {
		logger.Errorf("error to connect with database: %v", err)
		panic(err)
	}

	server := server.Setup()

	server.Run(fmt.Sprintf(":%s", conf.WebServerPort))
}
