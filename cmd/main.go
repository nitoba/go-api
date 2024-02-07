package main

import (
	"fmt"

	"github.com/nitoba/apis/configs"
	"github.com/nitoba/apis/internal/infra/database/gorm"
	"github.com/nitoba/apis/internal/infra/http/server"
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
