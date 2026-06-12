package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvCfg struct {
	Port       string
	APIAppName string
	DB_URL     string
}

func SetupEnvCfg() *EnvCfg {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	appName := os.Getenv("API_APP_NAME")
	if appName == "" {
		log.Fatal("API_APP_NAME environment variable is not set")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	return &EnvCfg{
		Port:       port,
		APIAppName: appName,
		DB_URL:     dbUrl,
	}
}
