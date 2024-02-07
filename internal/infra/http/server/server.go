package server

import (
	"github.com/labstack/echo"
	"github.com/nitoba/apis/internal/infra/http/server/routes"
)

func Setup() *echo.Echo {
	e := echo.New()
	routes.UsersRouter(e)
	return e
}
