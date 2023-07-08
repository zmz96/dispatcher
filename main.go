package main

import (
	"dispatcher/api"
	"dispatcher/config"
	"dispatcher/db"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
)

func initConfig(logger log.Logger) {
	// TODO get config path as flag
	config.Viper().AddConfigPath("./config")
	config.Viper().SetConfigName("default")
	config.Viper().SetConfigType("env")

	config.Viper().AutomaticEnv()
	if err := config.Viper().ReadInConfig(); err != nil {
		logger.Log("cannot load config file %s", err)
	}
	config.Parse()
}

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	initConfig(logger)

	db := db.Connect()

	router := gin.Default()
	api.RegisterRouterAPI(router, logger, db)
	router.Run(fmt.Sprintf("%s:%d", config.Config.HttpServer.Host, config.Config.HttpServer.Port))
}
