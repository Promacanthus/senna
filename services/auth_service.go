package services

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go/v4/request"

	"server-jwt/api/parameters"
	"server-jwt/core/authentication"
	"server-jwt/services/models"

	jwtv4 "github.com/dgrijalva/jwt-go/v4"
)

func Login(requestUser *models.User) (int, []byte) {
	authBackend := authentication.InitJWTAuthenticationBackend()

	if authBackend.Authenticate(requestUser) {
		token, err := authBackend.GenerateToken(requestUser.UUID)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			response, _ := json.Marshal(parameters.TokenAuthentication{Token: token})
			return http.StatusOK, response
		}
	}
	return http.StatusUnauthorized, []byte("")
}

func RefreshToken(requestUser *models.User) []byte {
	authBackend := authentication.InitJWTAuthenticationBackend()

	token, err := authBackend.GenerateToken(requestUser.UUID)
	if err != nil {
		panic(err)
	}

	response, err := json.Marshal(parameters.TokenAuthentication{Token: token})
	if err != nil {
		panic(err)
	}
	return response
}

func Logout(req *http.Request) error {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwtv4.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	if err != nil {
		return err
	}
	tokenString := req.Header.Get("Authorization")
	return authBackend.Logout(tokenString, tokenRequest)
}
