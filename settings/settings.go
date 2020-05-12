package settings

import (
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Settings struct {
	PrivateKeyPath     string
	PublicKeyPath      string
	JWTExpirationDelta int
}

var settings Settings = Settings{}

var env = "preproduction"

func Init() {

	viper.SetConfigName("pre")
	viper.SetConfigName("prod")
	viper.SetConfigName("tests")
	viper.SetConfigType("json")

	viper.AddConfigPath("./configuration")

	env = os.Getenv("GO_ENV")
	if env == "" {
		logrus.Warningln("Setting preproduction environment due to lack of GO_ENV value")
		env = "preproduction"
	}
	LoadConfiguration()
}

func LoadConfiguration() {

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorln(err)
	}

	err = viper.Unmarshal(&settings)
	if err != nil {
		logrus.Fatalln(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		logrus.Println("配置发生变更：", in.Name)
	})
}

func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}
