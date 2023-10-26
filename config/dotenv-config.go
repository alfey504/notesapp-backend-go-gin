package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var MONGODB_URI string = "MONGODB_URI"
var DB_NAME string = "DB_NAME"
var JWT_SECRET_KEY = "JWT_SECRET_KEY"

func GetEnvVariableFatal(key string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found")
	}

	dbName := os.Getenv(key)
	if dbName == "" {
		log.Fatal("DB_NAME is not properly set")
	}
	return dbName
}

func GetEnvVariableNonFatal(key string) (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}

	envValue := os.Getenv(key)
	if envValue == "" {
		return "", errors.New("No environment variable with key " + key)
	}
	return envValue, nil
}
