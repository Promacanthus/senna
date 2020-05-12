package main

import (
	"github.com/sirupsen/logrus"

	"server-jwt/routers"
	"server-jwt/settings"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	settings.Init()
	engine := routers.InitRouters()
	err := engine.Run(":5000")
	if err != nil {
		logrus.Error(err)
	}
}
