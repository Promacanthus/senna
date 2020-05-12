package routers

import "github.com/gin-gonic/gin"

func InitRouters() *gin.Engine {
	router := gin.Default()
	router = SetHelloRoutes(router)
	router = SetAuthenticationRoutes(router)
	return router
}
