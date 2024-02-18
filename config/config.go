package config

import "os"

type Config struct {
	//DATABASE
	DBDriver   string
	DBHost     string
	DBName     string
	DBPort     string
	DBUsername string
	DBPassword string
	ServerPort string
	ServerName string
}

func LoadConfig() Config {
	return Config{
		DBDriver:   os.Getenv("DB_DRIVER"),
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUsername: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		ServerPort: os.Getenv("API_PORT"),
		ServerName: os.Getenv("SERVER_NAME"),
	}
}
