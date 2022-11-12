package main

import (
	"fmt"
	"io"
	"log"
	"os"

	// "net/http"
	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	models "server/models"
)

func initLogToFile() {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(
		fmt.Sprintf("%s%s", os.Getenv("LOGGING_PATH"), os.Getenv("LOGGING_FILE_NAME")),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)

	if err != nil {
		log.Printf("Failed to Log into File: %s", err)
		return
	}

	log.SetOutput(io.MultiWriter(os.Stderr, file))
}

func initDB() (*gorm.DB, error) {

	var user string = os.Getenv("DB_USER")
	var password string = os.Getenv("DB_PASSWORD")
	var host string = os.Getenv("DB_HOST")
	var port string = os.Getenv("DB_PORT")
	var dbname string = os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func main() {

	// optional: enable Development environment

	if os.Getenv("DEVELOPMENT") == "true" {
		err := godotenv.Load("./config/.env")
		if err != nil {
			fmt.Printf("Error while loading .env file for DEVELOPMENT: %s", err)
		}
	}

	if os.Getenv("LOG_TO_FILE") == "true" {
		initLogToFile()
	}

	db, err := initDB()

	if err != nil {
		log.Fatalf("Error connecting to DB: %s", err)
	}

	if os.Getenv("MIGRATE") == "true" {
		db.AutoMigrate(&models.Request{})
	}
}
