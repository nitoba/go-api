package validations

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

func SendBadRequestValidation(data interface{}, c *gin.Context) bool {
	v := validate.Struct(data)

	if !v.Validate() {
		c.JSON(http.StatusBadRequest, gin.H{"message": v.Errors.Error()})
		return false
	}

	return true
}
