package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func LoadEnvVariables() (DbConfig, error) {
	err := godotenv.Load(".env") //metodo para cargar nuestras variables de un file.
	if err != nil {
		log.Fatalf("Error loading .env file")
		return DbConfig{}, err
	}
	return DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
	}, nil
}
