package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var Config = loadConfig()

func GetConfig() *Configuration {
	return Config
}

type Configuration struct {
	DbUsername string
	DbPassword string
	DbName     string
	DbHost     string
	DbPort     string
}

func loadConfig() *Configuration {

	return &Configuration{
		DbUsername: os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASS"),
		DbName:     os.Getenv("DB_NAME"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
	}
}
