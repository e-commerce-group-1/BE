package configs

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfigDev struct {
	Port     int `yaml:"port"`
	Database struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Address  string `yaml:"address"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

var lockDev = &sync.Mutex{}
var appConfigDev *AppConfig

func GetConfigDev() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfigDev() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 8000
	defaultConfig.Database.Driver = "mysql"
	defaultConfig.Database.Name = "test"
	defaultConfig.Database.Address = "localhost"
	defaultConfig.Database.Port = 3306
	defaultConfig.Database.Username = "admin"
	defaultConfig.Database.Password = "admin"

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./layered/configs/")
	if err := viper.ReadInConfig(); err != nil {
		log.Info("failed to open file")
		return &defaultConfig
	}

	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("failed to extract external config, use default value")
		return &defaultConfig
	}
	return &finalConfig
}
