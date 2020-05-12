package routers

import (
	"github.com/gin-gonic/gin"

	"server-jwt/controllers"
	"server-jwt/core/authentication"
)

func SetAuthenticationRoutes(router *gin.Engine) *gin.Engine {
	router.POST("/token-auth", controllers.Login)
	router.GET("/refresh-token-auth", authentication.RequireTokenAuthentication, controllers.RefreshToken)
	router.GET("/logout", authentication.RequireTokenAuthentication, controllers.Logout)
	return router
}
