package main

import (
	"fmt"
	"log"

	"github.com/Anton-Hudz/MovieList/app/repository"
	"github.com/Anton-Hudz/MovieList/app/transport"
	"github.com/Anton-Hudz/MovieList/app/usecase"
	"github.com/Anton-Hudz/MovieList/cfg"
	"github.com/spf13/viper"
)

func main() {

	repos := repository.NewRepository()
	usecase := usecase.NewUseCase(repos)
	handlers := transport.NewHandler(usecase)

	srv := new(transport.Server)

	////war vith .env
	// config, err := cfg.GetConfig()
	// if err != nil {
	// 	fmt.Printf("Failed to get config: %s", err)
	// }

	// if err := srv.Run(config.Server.Port, handlers.InitRoutes()); err != nil {
	// 	log.Fatal("Error occured while runnning server: %w", err.Error())
	// }

	////var with viper and .yml
	if err := cfg.InitConfig(); err != nil {
		fmt.Printf("Failed to get config: %s", err.Error())
	}

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal("Error occured while runnning server: %w", err.Error())
	}
}
