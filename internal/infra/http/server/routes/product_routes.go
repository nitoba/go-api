package routes

import (
	"github.com/gin-gonic/gin"
	usecases "github.com/nitoba/go-api/internal/domain/application/use_cases"
	"github.com/nitoba/go-api/internal/infra/database/gorm"
	"github.com/nitoba/go-api/internal/infra/database/gorm/repositories"
	"github.com/nitoba/go-api/internal/infra/http/controllers"
)

func ProductRouter(app *gin.Engine) {
	db := gorm.GetDB()

	productRepository := repositories.NewProductRepository(db)

	createProductUseCase := usecases.NewCreateProductUseCase(productRepository)
	createProductController := controllers.NewCreateProductController(createProductUseCase)
	router := app.Group("/products")
	{
		router.POST("/", createProductController.Handle)
	}
}
