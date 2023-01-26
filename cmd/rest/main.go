package main

import (
	"fmt"

	"github.com/Anton-Hudz/MovieList/app/repository"
	"github.com/Anton-Hudz/MovieList/app/transport"
	"github.com/Anton-Hudz/MovieList/app/usecase"
	"github.com/Anton-Hudz/MovieList/cfg"

	"github.com/Anton-Hudz/MovieList/logger"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	// easy "github.com/t-tomalak/logrus-easy-formatter"
)

func main() {

	////variant vith .env
	// config, err := cfg.GetConfig()
	// if err != nil {
	// 	fmt.Printf("Failed to get config: %s", err)
	// }
	// logger, err := logger.New(config.LogLevel)
	// if err != nil {
	// 	fmt.Printf("failed to create logger: %s", err)
	// }
	// db, err := repository.ConnectDB(config.DB)
	// if err != nil {
	// logger.Errorf("%+v", err)
	// 	fmt.Printf("Error connecting to database on host: %s, port: %s, with error: %s", config.DB.Host, config.DB.Port, err)
	// }
	// repos := repository.NewRepository(db)
	// usecase := usecase.NewUseCase(repos)
	// handlers := transport.NewHandler(usecase)

	// logger.Info("Connection to database successfully created")

	// srv := new(transport.Server)
	// logger.Infof("Server started on port: %v", config.Server.Port)
	// if err := srv.Run(config.Server.Port, handlers.InitRoutes()); err != nil {
	// 	logger.Fatalf("Error occured while runnning server: %w", err.Error())
	// }

	////variant with viper and .yml

	//
	// formatter := &logrus.TextFormatter{
	// 	FullTimestamp: true,
	// }
	// logrus.SetFormatter(formatter)

	// customFormatter := new(logrus.TextFormatter)
	// customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	// logrus.SetFormatter(customFormatter)
	// customFormatter.FullTimestamp = true

	logger.CustomLogger()

	// logrus := &logrus.Logger{
	// 	Out:   os.Stdout,
	// 	Level: logrus.InfoLevel,
	// 	Formatter: &easy.Formatter{
	// 		TimestampFormat: "2006-01-02 15:04:05",
	// 		LogFormat:       "[%lvl%]: %time% - message: %msg%\n",
	// 	},
	// }
	if err := cfg.InitConfig(); err != nil {
		fmt.Printf("Failed to get config: %s", err.Error())
	}

	config, err := cfg.GetViperConfig() // передавать сюда лог левел
	if err != nil {
		fmt.Printf("Failed to get Viper config: %s", err)
	}
	// logger, err := logger.New(config.LogLevel)
	// if err != nil {
	// 	fmt.Printf("failed to create logger: %s", err)
	// }
	db, err := repository.ConnectDB(config.DB)
	if err != nil {
		logrus.Errorf("%+v", err)
		fmt.Printf("Error connecting to database on host: %s, port: %s, with error: %s", viper.GetString("db.host"), viper.GetString("db.port"), err)

		return
	}

	repo := repository.NewRepository(db)
	usecase := usecase.NewUseCase(repo)
	handlers := transport.NewHandler(usecase)
	logrus.Info("Connection to database successfully created")

	srv := new(transport.Server)
	logrus.Infof("Server started on port: %v", config.Server.Port)
	if err := srv.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while runnning server: %w", err.Error())
	}
}
