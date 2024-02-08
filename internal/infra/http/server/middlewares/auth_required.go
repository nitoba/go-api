package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nitoba/go-api/internal/domain/application/cryptography"
	"github.com/nitoba/go-api/internal/domain/application/repositories"
)

func sendUnauthorizedResponse(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": msg,
	})
}

func AuthRequired(encrypter cryptography.Encrypter, userRepository repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		if authorization == "" || !strings.Contains(authorization, "Bearer") {
			sendUnauthorizedResponse(c, "Token Required")
			return
		}

		token := strings.TrimPrefix(authorization, "Bearer ")

		payload, err := encrypter.Verify(token)

		if err != nil {
			sendUnauthorizedResponse(c, "Invalid token")
			return
		}

		userID := payload["sub"].(string)

		user, err := userRepository.FindByID(userID)

		if err != nil || user == nil {
			sendUnauthorizedResponse(c, "Unauthorized")
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}
