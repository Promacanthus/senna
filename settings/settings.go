package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

var environments = map[string]string{
	"production":    "settings/prod.json",
	"preproduction": "settings/pre.json",
	"tests":         "settings/tests.json",
}

type Settings struct {
	PrivateKeyPath     string
	PublicKeyPath      string
	JWTExpirationDelta int
}

var settings Settings = Settings{}

var env = "preproduction"

func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		logrus.Warningln("Setting preproduction environment due to lack of GO_ENV value")
		env = "preproduction"
	}
	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		logrus.Fatalln("Error while reading config file", err)
	}
	settings = Settings{}
	err = json.Unmarshal(content, &settings)
	if err != nil {
		logrus.Fatalln("Error while parsing config file", err)
	}
}

func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}
