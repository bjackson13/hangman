package config

import (
	"github.com/joho/godotenv"
)

/*LoadEnvVariables read in our dot env file*/
func LoadEnvVariables() error {
	//load .env file
	err := godotenv.Load()
	return err
}