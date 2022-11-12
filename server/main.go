package main

import (
	"fmt"
	"os"

	// "net/http"
	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	models "server/models"
)

var DB *gorm.DB

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

	if os.Getenv("DEVELOPMENT") == "true" {

		err := godotenv.Load("./config/.env")
		if err != nil {
			fmt.Println("Error loading .env file")
		}

	}

	db, err := initDB()
	if err != nil {
		fmt.Printf("Error connecting to DB: %s", err)
	}

	if os.Getenv("MIGRATE") == "true" {
		db.AutoMigrate(&models.Request{})
	}
}
