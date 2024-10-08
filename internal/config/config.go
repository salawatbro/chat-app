package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	App App
	DB  DB
	JWT JWT
}

type DB struct {
	MongoDBURI   string
	DatabaseName string
}

type App struct {
	Port       string
	Env        string
	Name       string
	MaxRequest int
	Debug      bool
}

type JWT struct {
	Secret  string
	Expires time.Duration
}

func LoadConfig() (Config, error) {
	// Set the file name of the configuration file
	viper.SetConfigName("./")
	viper.SetConfigType("env") // Use .env format
	viper.AddConfigPath(".")   // Look for the config file in the working directory

	// Read the environment variables from .env file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return Config{
		App: App{
			Port:       viper.GetString("APP_PORT"),
			Env:        viper.GetString("APP_ENV"),
			Name:       viper.GetString("APP_NAME"),
			MaxRequest: viper.GetInt("APP_MAX_REQUEST"),
			Debug:      viper.GetBool("APP_DEBUG"),
		},
		DB: DB{
			MongoDBURI:   viper.GetString("MONGODB_URI"),
			DatabaseName: viper.GetString("DATABASE_NAME"),
		},
		JWT: JWT{
			Secret:  viper.GetString("JWT_SECRET"),
			Expires: viper.GetDuration("JWT_EXPIRES"),
		},
	}, nil
}
