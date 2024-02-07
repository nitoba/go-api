package routes

import (
	"github.com/labstack/echo"
	usecases "github.com/nitoba/apis/internal/domain/application/use_cases"
	"github.com/nitoba/apis/internal/infra/cryptography"
	"github.com/nitoba/apis/internal/infra/database/gorm"
	"github.com/nitoba/apis/internal/infra/database/gorm/repositories"
	"github.com/nitoba/apis/internal/infra/http/controllers"
)

func UsersRouter(e *echo.Echo) {
	bcryptHasher := cryptography.CreateBCryptHasher()
	userRepository := repositories.NewUserRepository(gorm.DBInstance)
	registerUserUseCase := usecases.CreateRegisterUseCase(userRepository, bcryptHasher)
	registerUserController := controllers.CreateRegisterUserController(registerUserUseCase)

	e.POST("/users", registerUserController.Handle)
}
