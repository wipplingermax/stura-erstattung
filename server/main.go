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
	var path string = "/var/log"
	var filename string = "stura-erstattung.log"

	if os.Getenv("LOG_PATH") != "" {
		path = os.Getenv("LOG_PATH")
	}

	if os.Getenv("LOG_FILE") != "" {
		filename = os.Getenv("LOG_FILE")
	}

	file, err := os.OpenFile(
		fmt.Sprintf("%s/%s", path, filename),
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

	var gormConfig gorm.Config = gorm.Config{}

	// configure Logging
	if os.Getenv("GORM_LOGGING") == "true" {

		// configure loglevel - default: Info
		var gormLogLevel logger.LogLevel

		switch os.Getenv("GORM_LOGLEVEL") {
		case "Silent":
			gormLogLevel = logger.Silent
		case "Error":
			gormLogLevel = logger.Error
		case "Warn":
			gormLogLevel = logger.Warn
		case "Info":
			gormLogLevel = logger.Info
		default:
			gormLogLevel = logger.Info
		}

		// configure file logging - default path: /var/log/stura-erstattung.log
		gormLogPaths := io.Writer(os.Stdout)

		if os.Getenv("GORM_LOG_TO_FILE") == "true" {

			var path string = "/var/log"
			var filename string = "stura-erstattung-db.log"

			if os.Getenv("GORM_LOG_PATH") != "" {
				path = os.Getenv("GORM_LOG_PATH")
			}

			if os.Getenv("GORM_LOG_FILE") != "" {
				filename = os.Getenv("GORM_LOG_FILE")
			}

			file, err := os.OpenFile(
				fmt.Sprintf("%s/%s", path, filename),
				os.O_APPEND|os.O_CREATE|os.O_WRONLY,
				0666)

			if err != nil {
				log.Printf("Failed to Log into File: %s", err)
			}

			gormLogPaths = io.MultiWriter(os.Stdout, file)
		}

		gormLogger := logger.New(
			log.New(gormLogPaths, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				LogLevel:                  gormLogLevel, // Log level
				IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,        // Disable color
			},
		)

		gormConfig = gorm.Config{Logger: gormLogger}
	}

	return gorm.Open(postgres.Open(dsn), &gormConfig)
}

func main() {

	// enable Development environment
	if os.Getenv("DEVELOPMENT") == "true" {
		err := godotenv.Load("./config/.env")
		if err != nil {
			fmt.Println("Error while loading .env file for DEVELOPMENT")
		}
	}

	// enable additional logging to file
	if os.Getenv("LOG_TO_FILE") == "true" {
		initLogToFile()
	}

	// initialize database connection
	db, err := initDB()

	if err != nil {
		log.Fatal("Error while initializing DB")
	}

	// optional: migrate database
	if os.Getenv("DB_MIGRATE") == "true" {
		db.AutoMigrate(&models.Request{})
	}
}
