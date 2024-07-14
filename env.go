package config

import (
	`github.com/joho/godotenv`
)

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return errUnableToLoadDotEnvFile()
	}

	return nil
}
