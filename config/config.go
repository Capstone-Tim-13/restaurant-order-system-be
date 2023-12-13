package config

import (
	"os"
	"strconv"
)

type DatabaseConfig struct {
	DB_USER         string
	DB_PASS         string
	DB_HOST         string
	DB_PORT         int
	DB_NAME         string
	CDN_CLOUD_NAME  string
	CDN_API_KEY     string
	CDN_API_SECRET  string
	CDN_FOLDER_NAME string
	SERVER_KEY_MT   string
}

type ServerConfig struct {
	SERVER_PORT string
}

func LoadDBConfig() DatabaseConfig {
	//godotenv.Load(".env")

	DB_PORT, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		panic(err)
	}

	return DatabaseConfig{
		DB_USER:         os.Getenv("DB_USER"),
		DB_PASS:         os.Getenv("DB_PASS"),
		DB_HOST:         os.Getenv("DB_HOST"),
		DB_PORT:         DB_PORT,
		DB_NAME:         os.Getenv("DB_NAME"),
		CDN_CLOUD_NAME:  os.Getenv("CDN_CLOUD_NAME"),
		CDN_API_KEY:     os.Getenv("CDN_API_KEY"),
		CDN_API_SECRET:  os.Getenv("CDN_API_SECRET"),
		CDN_FOLDER_NAME: os.Getenv("CDN_FOLDER_NAME"),
		SERVER_KEY_MT:   os.Getenv("SERVER_KEY_MT"),
	}
}
