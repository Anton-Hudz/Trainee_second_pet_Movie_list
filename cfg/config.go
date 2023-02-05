package cfg

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	lenOfLines  = 2
	envFileName = ".env"
)

type Options struct {
	SigningKey string
	LogLevel   string
	Server     Server
	DB         DB
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

//if you want to use just .env file to conf.
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
		SigningKey: os.Getenv("SIGNING_KEY"),
		LogLevel:   os.Getenv("LOG_LEVEL"),
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

//for viper:
func GetViperConfig() (Options, error) {
	if err := godotenv.Load(); err != nil {
		return Options{}, errors.New("error loading env file")
	}

	opt := Options{
		SigningKey: viper.GetString("token.signingKey"),
		LogLevel:   viper.GetString("logger.loglevel"),
		Server: Server{
			Host: viper.GetString("server.host"),
			Port: viper.GetString("server.port"),
		},
		DB: DB{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			Password: os.Getenv("DB_PASSWORD"), //this is safety data and you can't add password to public repo
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
		},
	}

	return opt, nil
}

func InitConfig() error {
	viper.AddConfigPath("cfg")    //name config dir
	viper.SetConfigName("config") //name config file
	return viper.ReadInConfig()
}
