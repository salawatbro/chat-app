package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App App
	DB  DB
}

type DB struct {
	MongoDBURI   string
	DatabaseName string
}

type App struct {
	Port string
	Env  string
	Name string
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
			Port: viper.GetString("APP_PORT"),
			Env:  viper.GetString("APP_ENV"),
			Name: viper.GetString("APP_NAME"),
		},
		DB: DB{
			MongoDBURI:   viper.GetString("MONGODB_URI"),
			DatabaseName: viper.GetString("DATABASE_NAME"),
		},
	}, nil
}
