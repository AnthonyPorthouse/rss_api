package main

import "os"

import _ "github.com/joho/godotenv/autoload"

type Configuration struct {
	DbUsername string
	DbPassword string
	DbName     string
	DbHost     string
	DbPort     string
}

func loadConfig() Configuration {

	return Configuration{
		DbUsername: os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASS"),
		DbName:     os.Getenv("DB_NAME"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
	}
}

var Config = loadConfig()
