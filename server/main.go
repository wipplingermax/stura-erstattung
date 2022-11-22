package main

import (
	"fmt"

	// "net/http"
	config "server/config"
	models "server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	// load environment, config logging
	config.Init()

	// initialize database connection
	db := config.InitDB()

	// initialize Router

	// test usage
	result := db.First(&models.Request{})
	fmt.Println(result.Error)

	router := gin.Default()

	router.POST("/formRequest", func(c *gin.Context) {
		fmt.Println("Request kam an")
	})
	router.Run()
}
