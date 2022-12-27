package main

import (
	"fmt"
	"log"

	"github.com/Anton-Hudz/MovieList/app/repository"
	"github.com/Anton-Hudz/MovieList/app/transport"
	"github.com/Anton-Hudz/MovieList/app/usecase"
	"github.com/Anton-Hudz/MovieList/cfg"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	////variant vith .env
	// config, err := cfg.GetConfig()
	// if err != nil {
	// 	fmt.Printf("Failed to get config: %s", err)
	// }
	// db, err := repository.ConnectDB(config.DB)
	// if err != nil {
	// 	fmt.Printf("Error connecting to database on host: %s, port: %s, with error: %s", config.DB.Host, config.DB.Port, err)
	// }
	// repos := repository.NewRepository(db)
	// usecase := usecase.NewUseCase(repos)
	// handlers := transport.NewHandler(usecase)

	// log.Println("Connection to database successfully created")

	// srv := new(transport.Server)
	// if err := srv.Run(config.Server.Port, handlers.InitRoutes()); err != nil {
	// 	log.Fatal("Error occured while runnning server: %w", err.Error())
	// }

	////variant with viper and .yml
	if err := cfg.InitConfig(); err != nil {
		fmt.Printf("Failed to get config: %s", err.Error())
	}

	config, err := cfg.GetViperConfig()
	if err != nil {
		fmt.Printf("Failed to get Viper config: %s", err)
	}
	db, err := repository.ConnectDB(config.DB)
	if err != nil {
		fmt.Printf("Error connecting to database on host: %s, port: %s, with error: %s", config.DB.Host, config.DB.Port, err)
		return
	}

	repo := repository.NewRepository(db)
	usecase := usecase.NewUseCase(repo)
	handlers := transport.NewHandler(usecase)

	log.Println("Connection to database successfully created")

	srv := new(transport.Server)
	if err := srv.Run(viper.GetString(config.Server.Port), handlers.InitRoutes()); err != nil {
		log.Fatal("Error occured while runnning server: %w", err.Error())
	}
}
