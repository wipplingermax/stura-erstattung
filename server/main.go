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
	"gorm.io/gorm/logger"

	models "server/models"
)

func initLogToFile() {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(
		fmt.Sprintf("%s/%s", os.Getenv("LOGGING_PATH"), os.Getenv("LOGGING_FILE")),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)

	if err != nil {
		log.Printf("Failed to Log into File: %s", err)
		return
	}

	log.SetOutput(io.MultiWriter(os.Stderr, file))
}

func initDB() (*gorm.DB, error) {

	// configure Connection
	var user string = os.Getenv("DB_USER")
	var password string = os.Getenv("DB_PASSWORD")
	var host string = os.Getenv("DB_HOST")
	var port string = os.Getenv("DB_PORT")
	var dbname string = os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	// configure Logging
	var gormConfig gorm.Config
	var gormLogLevel logger.LogLevel

	if os.Getenv("GORM_LOGGING") == "true" {
		switch os.Getenv("GORM_LOGLEVEL") {
		case "Silent":
			gormLogLevel = logger.Silent
		case "Error":
			gormLogLevel = logger.Error
		case "Warn":
			gormLogLevel = logger.Warn
		case "Info":
			gormLogLevel = logger.Info
		}
		gormConfig = gorm.Config{Logger: logger.Default.LogMode(gormLogLevel)}
	} else {
		gormConfig = gorm.Config{}
	}

	return gorm.Open(postgres.Open(dsn), &gormConfig)
}

func main() {

	// enable Development environment
	if os.Getenv("DEVELOPMENT") == "true" {
		err := godotenv.Load("./config/.env")
		if err != nil {
			fmt.Printf("Error while loading .env file for DEVELOPMENT: %s", err)
		}
	}

	// enable additional logging to file
	if os.Getenv("LOG_TO_FILE") == "true" {
		initLogToFile()
	}

	// initialize database connection
	db, err := initDB()

	if err != nil {
		log.Fatalf("Error while initializing DB: %s", err)
	}

	// optional: migrate database
	if os.Getenv("DB_MIGRATE") == "true" {
		db.AutoMigrate(&models.Request{})
	}
}
