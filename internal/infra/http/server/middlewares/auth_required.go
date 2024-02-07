package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nitoba/go-api/internal/infra/cryptography"
	"github.com/nitoba/go-api/internal/infra/database/gorm"
	"github.com/nitoba/go-api/internal/infra/database/gorm/repositories"
)

func AuthRequired() gin.HandlerFunc {
	db := gorm.GetDB()
	jwtEncrypter := cryptography.NewJWTEncrypter()
	userRepository := repositories.NewUserRepository(db)
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		if authorization == "" || !strings.Contains(authorization, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		token := strings.TrimPrefix(authorization, "Bearer ")

		payload, err := jwtEncrypter.Verify(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		userID := payload["sub"].(string)

		user, err := userRepository.FindByID(userID)

		if err != nil || user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}
