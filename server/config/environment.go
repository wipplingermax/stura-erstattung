package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error while loading .env file for DEVELOPMENT")
	}
}
