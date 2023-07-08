package config

import (
	"log"

	"github.com/spf13/viper"
)

var viperInstance = viper.New()
var Config Configs

type Configs struct {
	HttpServer struct {
		Port        uint   `mapstructure:"port"`
		Host        string `mapstructure:"host"`
		LogFileName string `mapstructure:"log_file_name"`
		LogLevel    int    `mapstructure:"log_level"`
	} `mapstructure:"http_server"`
	Database struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"pass"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		DBName   string `mapstructure:"name"`
		SSLMode  string `mapstructure:"ssl_mode"`
		TimeZone string `mapstructure:"time_zone"`
	} `mapstructure:"db"`
}

func Parse() Configs {
	if err := viperInstance.Unmarshal(&Config); err != nil {
		log.Fatalf("cannot read config file %s", err)
	}
	return Config
}

func Viper() *viper.Viper {
	return viperInstance
}
