package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/janrockdev/golang-mongodb-rest/models"
)

func AppRecovery() func(c *gin.Context, recovered interface{}) {
	return func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			models.SendErrorResponse(c, http.StatusInternalServerError, err)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false}) // recovery failed
	}
}
