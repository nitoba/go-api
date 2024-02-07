package routes

import (
	"github.com/gin-gonic/gin"
	usecases "github.com/nitoba/go-api/internal/domain/application/use_cases"
	"github.com/nitoba/go-api/internal/infra/cryptography"
	"github.com/nitoba/go-api/internal/infra/database/gorm"
	"github.com/nitoba/go-api/internal/infra/database/gorm/repositories"
	"github.com/nitoba/go-api/internal/infra/http/controllers"
)

func AuthRouter(app *gin.Engine) {
	db := gorm.GetDB()
	bcryptHasher := cryptography.CreateBCryptHasher()
	userRepository := repositories.NewUserRepository(db)
	registerUserUseCase := usecases.CreateRegisterUseCase(userRepository, bcryptHasher)
	registerUserController := controllers.CreateRegisterUserController(registerUserUseCase)

	router := app.Group("/auth")
	{
		router.POST("/register", registerUserController.Handle)
	}
}
