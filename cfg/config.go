package cfg

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

const (
	lenOfLines  = 2
	envFileName = ".env"
)

type Options struct {
	Server Server
	DB     DB
}

type Server struct {
	Host string
	Port string
}

type DB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func loadEnvVar() error {
	f, err := os.Open(envFileName)
	if err != nil {
		return fmt.Errorf("error occured while opening %s file: %w", envFileName, err)
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Printf("%s", err)
		}
	}()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error occured while scanning %s file: %w", envFileName, err)
	}

	for _, l := range lines {
		pair := strings.Split(l, "=")
		if len(pair) != lenOfLines {
			return errors.New("not enough data for the configuration at the config file")
		}
		os.Setenv(pair[0], pair[1])
	}

	return nil
}

func GetConfig() (Options, error) {
	if err := loadEnvVar(); err != nil {
		return Options{}, err
	}

	opt := Options{
		Server: Server{
			Host: os.Getenv("SERV_HOST"),
			Port: os.Getenv("SERV_PORT"),
		},
		DB: DB{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
	}

	return opt, nil
}

func InitConfig() error {
	viper.AddConfigPath("cfg")    //name config dir
	viper.SetConfigName("config") //name config file
	return viper.ReadInConfig()
}
