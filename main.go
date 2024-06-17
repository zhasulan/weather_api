package main

import (
	"go.uber.org/zap"
	"net/http"
	"weather_api/internal/config"
	"weather_api/internal/database"
	"weather_api/internal/logger"
	"weather_api/internal/repository"
	"weather_api/internal/router"
	"weather_api/internal/services"
)

func main() {
	// config
	config.InitConfig(config.GetConfigPath())

	// services
	services.WEATHER_API = services.NewWeatherService()

	// logger
	logger.InitLogger()

	// db
	db, err := database.Connection(config.Config.DB)
	if err != nil {
		// todo logging
		return
	}
	defer db.Close()

	// check db
	err = db.Ping()
	if err != nil {
		// todo logging
		return
	}

	// env
	{
		env := config.Env{
			DB: db,
			Repository: &repository.IRepository{
				DB: db,
			},
		}
		config.SetEnv(env)
	}

	// server
	httpServer := http.Server{
		Addr:    config.Config.App.Host + ":" + config.Config.App.Port,
		Handler: router.NewRouter(),
	}

	zap.S().Panic(httpServer.ListenAndServe())

	// todo graceful shutdown
}
