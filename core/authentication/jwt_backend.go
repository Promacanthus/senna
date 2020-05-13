package authentication

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"time"

	jwtv4 "github.com/dgrijalva/jwt-go/v4"
	"github.com/pborman/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"server-jwt/core/redis"
	"server-jwt/services/models"
	"server-jwt/settings"
)

const expireOffset = 3600

type JWTAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

var authBackendInstance *JWTAuthenticationBackend = nil

func InitJWTAuthenticationBackend() *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			privateKey: getPrivateKey(),
			PublicKey:  getPublicKey(),
		}
	}

	return authBackendInstance
}

func (j *JWTAuthenticationBackend) GenerateToken(userUUID string) (string, error) {

	// creat a new token object with a specified signing method and the claims
	token := jwtv4.New(jwtv4.SigningMethodRS512)

	token.Claims = jwtv4.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(settings.Get().JWTExpirationDelta)).Unix(),
		"iat": time.Now().Unix(),
		"sub": userUUID,
	}

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(j.privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWTAuthenticationBackend) Authenticate(user *models.User) bool {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("testing"), 10)

	testUser := &models.User{
		UUID:     uuid.New(),
		Username: "haku",
		Password: string(hashPassword),
	}

	return user.Username == testUser.Username &&
		bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)) == nil
}

func (j *JWTAuthenticationBackend) getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainder := time.Until(tm)
		if remainder > 0 {
			return int(remainder.Seconds() + expireOffset)
		}
	}
	return expireOffset
}

func (j *JWTAuthenticationBackend) Logout(tokenString string, token *jwtv4.Token) error {
	redisConn := redis.Connect()
	return redisConn.SetValue(tokenString, tokenString, j.getTokenRemainingValidity(token.Claims.(jwtv4.MapClaims)["exp"]))
}

func (j *JWTAuthenticationBackend) IsInBlacklist(token string) bool {
	redisConn := redis.Connect()
	redisToken, _ := redisConn.GetValue(token)
	return redisToken != nil
}

func getPrivateKey() *rsa.PrivateKey {

	privateKeyFile, err := os.Open(settings.Get().PrivateKeyPath)
	if err != nil {
		panic(err)
	}

	pemFileInfo, _ := privateKeyFile.Stat()
	var size = pemFileInfo.Size()
	pemBytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pemBytes)
	if err != nil {
		logrus.Fatalln(err)
	}

	data, _ := pem.Decode(pemBytes)

	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		panic(err)
	}
	return privateKeyImported
}

func getPublicKey() *rsa.PublicKey {
	publicKeyFile, err := os.Open(settings.Get().PublicKeyPath)
	if err != nil {
		panic(err)
	}

	pemFileInfo, _ := publicKeyFile.Stat()
	var size = pemFileInfo.Size()
	pemBytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pemBytes)
	if err != nil {
		logrus.Fatalln(err)
	}

	data, _ := pem.Decode(pemBytes)

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)
	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}
