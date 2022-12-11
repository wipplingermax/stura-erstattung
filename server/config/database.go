package config

import (
	"fmt"
	"io"
	"log"
	"os"

	loggable "github.com/daqingshu/gorm-loggable"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"server/models"
)

func automigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Request{})
	db.AutoMigrate(&models.Refund{})
	db.AutoMigrate(&models.Transcation{})
	db.AutoMigrate(&models.User{})
}

func initGormConfig() gorm.Config {

	// configure loglevel
	// default: Info
	var gormLogLevel logger.LogLevel = logger.Info

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

	// configure file logging
	// default path: /var/log/stura-erstattung.log
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
			Colorful:                  true,         // Disable color
		},
	)
	return gorm.Config{Logger: gormLogger}
}

func InitDB() *gorm.DB {

	// configure Connection
	var user string = os.Getenv("DB_USER")
	var password string = os.Getenv("DB_PASSWORD")
	var host string = os.Getenv("DB_HOST")
	var port string = os.Getenv("DB_PORT")
	var dbname string = os.Getenv("DB_NAME")

	var dsn string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	// configure GORM
	var gormConfig gorm.Config = initGormConfig()

	// initialize Connection
	db, err := gorm.Open(postgres.Open(dsn), &gormConfig)

	if err != nil {
		log.Fatal("Error while initializing DB")
	}

	// optional: migrate database
	if os.Getenv("DB_MIGRATE") == "true" {
		automigrate(db)
	}

	_, lerr := loggable.Register(db)
	if lerr != nil {
		panic("couldn't attach logging function to database")
	}

	return db
}
