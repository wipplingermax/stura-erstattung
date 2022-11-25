package main

import (
	"github.com/gin-gonic/gin"

	config "server/config"
	controller "server/controller"
)

var router *gin.Engine

func init() {

	// load environment, config logging
	config.Init()

	// initialize database connection
	config.InitDB()

	// initialize Router

}

func main() {

	router := gin.Default()
	router.POST("/v1/request", controller.CreateRequest)
	router.Run()

}
