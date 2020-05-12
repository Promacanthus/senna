package authentication

import (
	"fmt"
	"net/http"

	jwtv4 "github.com/dgrijalva/jwt-go/v4"
	"github.com/dgrijalva/jwt-go/v4/request"
	"github.com/gin-gonic/gin"
)

func RequireTokenAuthentication(ctx *gin.Context) {
	authBackend := InitJWTAuthenticationBackend()

	token, err := request.ParseFromRequest(ctx.Request, request.OAuth2Extractor, func(token *jwtv4.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtv4.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		} else {
			return authBackend.PublicKey, nil
		}
	})

	if err == nil && token.Valid && !authBackend.IsInBlacklist(ctx.Request.Header.Get("Authorization")) {
		ctx.Next()
	} else {
		ctx.Writer.WriteHeader(http.StatusUnauthorized)
	}
}
