package config

import "os"

var (
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	APP_PORT    string
	SECRET_KEY  string
)

func InitConfigEnv() {
	appPortEnv := os.Getenv("APP_PORT")
	if appPortEnv != "" {
		APP_PORT = appPortEnv
	}

	dbHostEnv := os.Getenv("DB_HOST")
	if dbHostEnv != "" {
		DB_HOST = dbHostEnv
	}

	dbNameEnv := os.Getenv("DB_NAME")
	if dbNameEnv != "" {
		DB_NAME = dbNameEnv
	}

	dbPasswordEnv := os.Getenv("DB_PASSWORD")
	if dbPasswordEnv != "" {
		DB_PASSWORD = dbPasswordEnv
	}

	dbPortEnv := os.Getenv("DB_PORT")
	if dbPortEnv != "" {
		DB_PORT = dbPortEnv
	}

	dbUserEnv := os.Getenv("DB_USER")
	if dbUserEnv != "" {
		DB_USER = dbUserEnv
	}

	secretKeyEnv := os.Getenv("SECRET_KEY")
	if secretKeyEnv != "" {
		SECRET_KEY = secretKeyEnv
	}
}
