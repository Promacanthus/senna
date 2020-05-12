package controllers

import (
	"github.com/gin-gonic/gin"
)

func HelloController(ctx *gin.Context) {
	ctx.Writer.Write([]byte("Hello,World!"))
}
