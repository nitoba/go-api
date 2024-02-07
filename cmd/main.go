package main

import (
	"fmt"

	"github.com/nitoba/go-api/configs"
	"github.com/nitoba/go-api/internal/infra/database/gorm"
	"github.com/nitoba/go-api/internal/infra/http/server"
)

func main() {
	logger := configs.GetLogger("main")
	conf, err := configs.LoadConfig()

	if err != nil {
		logger.Errorf("error loading config: %v", err)
		panic(err)
	}

	err = gorm.Connect()

	if err != nil {
		panic(err)
	}

	server := server.Setup()

	server.Run(fmt.Sprintf(":%s", conf.WebServerPort))
}
