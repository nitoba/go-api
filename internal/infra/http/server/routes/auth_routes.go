package routes

import (
	"github.com/gin-gonic/gin"
	usecases "github.com/nitoba/go-api/internal/domain/application/use_cases"
	"github.com/nitoba/go-api/internal/infra/cryptography"
	"github.com/nitoba/go-api/internal/infra/database/prisma"
	"github.com/nitoba/go-api/internal/infra/database/prisma/repositories"
	"github.com/nitoba/go-api/internal/infra/http/controllers"
)

func AuthRouter(app *gin.Engine) {
	db := prisma.GetDB()

	jwtEncrypter := cryptography.NewJWTEncrypter()
	bcryptHasher := cryptography.CreateBCryptHasher()
	userRepository := repositories.NewUserRepositoryPrisma(db)

	registerUserUseCase := usecases.CreateRegisterUseCase(userRepository, bcryptHasher)
	authenticateUserUseCase := usecases.NewAuthenticate(userRepository, bcryptHasher, jwtEncrypter)

	registerUserController := controllers.CreateRegisterUserController(registerUserUseCase)
	authenticateUserController := controllers.NewAuthenticateUserController(authenticateUserUseCase)

	router := app.Group("/auth")
	{
		router.POST("/register", registerUserController.Handle)
		router.POST("/authenticate", authenticateUserController.Handle)
	}
}
