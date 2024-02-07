package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nitoba/go-api/internal/infra/http/server/routes"
)

func Setup() *gin.Engine {
	r := gin.Default()
	routes.AuthRouter(r)
	routes.ProductRouter(r)
	return r
}
