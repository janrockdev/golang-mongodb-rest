package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/janrockdev/golang-mongodb-rest/controllers"
)

func PingRoute(router *gin.RouterGroup) {
	auth := router.Group("/ping")
	{
		auth.GET(
			"",
			controllers.Ping,
		)
	}
}
