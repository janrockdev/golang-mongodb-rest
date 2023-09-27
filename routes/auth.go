package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/janrockdev/golang-mongodb-rest/controllers"
	"github.com/janrockdev/golang-mongodb-rest/middlewares/validators"
)

func AuthRoute(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST(
			"/register",
			validators.RegisterValidator(),
			controllers.Register,
		)

		auth.POST(
			"/login",
			validators.LoginValidator(),
			controllers.Login,
		)

		auth.POST(
			"/refresh",
			validators.RefreshValidator(),
			controllers.Refresh,
		)
	}
}
