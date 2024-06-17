package config

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

func InitConfig() (*AppConfig, *DBConfig) {
	db := LoadDB()
	app := LoadAPP()

	if db == nil || app == nil {
		logrus.Fatal("Config : cannot start program, failed to load configuration")
		return nil, nil
	}
	return app, db
}

type AppConfig struct {
	APP_PORT   int
	SECRET_KEY string
}

func LoadAPP() *AppConfig {
	var result = new(AppConfig)

	if v, found := os.LookupEnv("APP_PORT"); found {
		port, err := strconv.Atoi(v)
		if err != nil {
			logrus.Error("Config: invalid port value,", err.Error())
			return nil
		}
		result.APP_PORT = port
	}
	return result
}
