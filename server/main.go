package main

import (
	"fmt"

	// "net/http"
	// "github.com/gin-gonic/gin"
	config "server/config"
	models "server/models"
)

func main() {

	// load environment, config logging
	config.Init()

	// initialize database connection
	db := config.InitDB()

	// initialize Router

	// test usage
	result := db.First(&models.Request{})
	fmt.Println(result.Error)

}
