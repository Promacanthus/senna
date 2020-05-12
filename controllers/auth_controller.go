package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"

	"server-jwt/services"
	"server-jwt/services/models"
)

func Login(ctx *gin.Context) {
	requestUser := &models.User{}
	err := json.NewDecoder(ctx.Request.Body).Decode(requestUser)
	if err != nil {
		logrus.Fatal(err)
	}

	responseStatus, token := services.Login(requestUser)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(responseStatus)
	ctx.Writer.Write(token)

}

func RefreshToken(ctx *gin.Context) {
	requestUser := &models.User{}
	json.NewDecoder(ctx.Request.Body).Decode(requestUser)

	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.Write(services.RefreshToken(requestUser))
}

func Logout(ctx *gin.Context) {
	err := services.Logout(ctx.Request)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
	} else {
		ctx.Writer.WriteHeader(http.StatusOK)
	}
}
