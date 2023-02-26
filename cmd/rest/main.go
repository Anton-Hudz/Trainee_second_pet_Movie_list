package main

import (
	"log"

	"github.com/Anton-Hudz/MovieList/app/repository"
	"github.com/Anton-Hudz/MovieList/app/transport"
	"github.com/Anton-Hudz/MovieList/app/usecase"
	"github.com/Anton-Hudz/MovieList/cfg"

	"github.com/Anton-Hudz/MovieList/logger"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	////variant vith .env
	// config, err := cfg.GetConfig()
	// if err != nil {
	// 	log.Printf("Failed to get config: %s", err)

	// 	return
	// }

	// if err := logger.CustomLogger(config.LogLevel); err != nil {
	// 	log.Printf("Failed to create logger: %s", err)

	// 	return
	// }

	// db, err := repository.ConnectDB(config.DB)
	// if err != nil {
	// 	logrus.Errorf("Error connecting to database on host: %s, port: %s, with error: %s", config.DB.Host, config.DB.Port, err)

	// 	return
	// }

	// repo := repository.NewRepo(db)

	// userUsecase := usecase.NewAuthUser(repo)
	// filmUsecase := usecase.NewFilmService(repo)

	// handlers := transport.NewHandler(userUsecase, filmUsecase)

	// logrus.Info("Connection to database successfully created")

	// srv := new(transport.Server)
	// logrus.Infof("Server started on port: %v", config.Server.Port)
	// if err := srv.Run(config.Server.Port, handlers.InitRoutes()); err != nil {
	// 	logrus.Fatalf("Error occured while runnning server: %w", err.Error())

	// 	return
	// }

	////variant with viper and .yml
	if err := cfg.InitConfig(); err != nil {
		log.Printf("Failed to get config: %s", err.Error())

		return
	}

	config, err := cfg.GetViperConfig()
	if err != nil {
		log.Printf("Failed to get Viper config: %s", err)

		return
	}

	if err := logger.CustomLogger(config.LogLevel); err != nil {
		log.Printf("Failed to create logger: %s", err)

		return
	}

	db, err := repository.ConnectDB(config.DB)
	if err != nil {
		logrus.Errorf("Error connecting to database on host: %s, port: %s, with error: %s", viper.GetString("db.host"), viper.GetString("db.port"), err)

		return
	}

	repo := repository.NewRepo(db)

	userUsecase := usecase.NewAuthUser(repo)
	filmUsecase := usecase.NewFilmService(repo)

	handlers := transport.NewHandler(userUsecase, filmUsecase)
	logrus.Info("Connection to database successfully created")

	srv := new(transport.Server)
	logrus.Infof("Server started on port: %v", config.Server.Port)
	if err := srv.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while runnning server: %w", err.Error())

		return
	}
}
