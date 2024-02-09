package routes

import (
	"github.com/gin-gonic/gin"
	usecases "github.com/nitoba/go-api/internal/domain/application/use_cases"
	"github.com/nitoba/go-api/internal/infra/cryptography"
	"github.com/nitoba/go-api/internal/infra/database/prisma"
	"github.com/nitoba/go-api/internal/infra/database/prisma/repositories"
	"github.com/nitoba/go-api/internal/infra/http/controllers"
	"github.com/nitoba/go-api/internal/infra/http/server/middlewares"
)

func ProductRouter(app *gin.Engine) {
	db := prisma.GetDB()
	jwtEncrypter := cryptography.NewJWTEncrypter()
	userRepository := repositories.NewUserRepositoryPrisma(db)
	productRepository := repositories.NewProductRepositoryPrisma(db)

	createProductUseCase := usecases.NewCreateProductUseCase(productRepository)
	createProductController := controllers.NewCreateProductController(createProductUseCase)

	router := app.Group("/products")
	router.Use(middlewares.AuthRequired(jwtEncrypter, userRepository))
	{
		router.POST("/", createProductController.Handle)
	}
}
