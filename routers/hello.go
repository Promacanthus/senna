package routers

import (
	"github.com/gin-gonic/gin"

	"server-jwt/controllers"
	"server-jwt/core/authentication"
)

func SetHelloRoutes(router *gin.Engine) *gin.Engine {
	router.GET("/test/hello", authentication.RequireTokenAuthentication, controllers.HelloController)
	return router
}
