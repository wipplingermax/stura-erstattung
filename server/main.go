package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	config "server/config"
	controller "server/controller"
	models "server/models"
)

var DB *gorm.DB

func main() {

	// load environment, config logging
	config.Init()

	// initialize database connection
	config.InitDB()
	DB = config.DB

	// initialize Router
	router := gin.Default()

	DB.AutoMigrate(&models.Request{})

	router.POST("/v1/request", controller.CreateRequest)
	router.Run()
}
