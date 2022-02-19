package configs

import (
	"fmt"
	"os"
	"sync"
)

type AppConfig struct {
	Port     int    `yaml:"port"`
	Driver   string `yaml:"driver"`
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	DB_Port  int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

// func initConfig() *AppConfig {
// 	var defaultConfig AppConfig
// 	defaultConfig.Port = 8000
// 	defaultConfig.Driver = "mysql"
// 	defaultConfig.Name = "test"
// 	defaultConfig.Address = "localhost"
// 	defaultConfig.DB_Port = 3306
// 	defaultConfig.Username = "admin"
// 	defaultConfig.Password = "admin"

// 	viper.SetConfigType("yaml")
// 	viper.SetConfigName("config")
// 	viper.AddConfigPath("./configs/")
// 	if err := viper.ReadInConfig(); err != nil {
// 		log.Info("failed to open file")
// 		return &defaultConfig
// 	}

// 	var finalConfig AppConfig
// 	err := viper.Unmarshal(&finalConfig)
// 	if err != nil {
// 		log.Info("failed to extract external config, use default value")
// 		return &defaultConfig
// 	}
// 	return &finalConfig
// }
func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 8000
	defaultConfig.Driver = getEnv("DRIVER", "mysql")
	defaultConfig.Name = getEnv("NAME", "test")
	defaultConfig.Address = getEnv("ADDRESS", "localhost")
	defaultConfig.DB_Port = 3306
	defaultConfig.Username = getEnv("USERNAME", "root")
	defaultConfig.Password = getEnv("PASSWORD", "mysqlku")

	fmt.Println(defaultConfig)

	return &defaultConfig
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "KCK" {
		fmt.Println(value)
		return value
	}
	return fallback
}
