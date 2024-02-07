package main

import (
	"fmt"

	"github.com/nitoba/go-api/configs"
	"github.com/nitoba/go-api/internal/infra/database/gorm"
	"github.com/nitoba/go-api/internal/infra/http/server"
)

func main() {
	conf, err := configs.LoadConfig()

	if err != nil {
		panic(err)
	}

	err = gorm.NewConnectionDB()

	if err != nil {
		panic(err)
	}

	server := server.Setup()

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", conf.WebServerPort)))
}
