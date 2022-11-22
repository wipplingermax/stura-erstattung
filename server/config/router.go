package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {

	r := gin.Default()

	// configure loglevel
	// default: Debug
	mode := gin.DebugMode

	switch os.Getenv("GIN_MODE") {
	case "Release":
		mode = gin.ReleaseMode
	case "Test":
		mode = gin.TestMode
	case "Debug":
		mode = gin.DebugMode
	}

	gin.SetMode(mode)

	//adding routes
	v1 := r.Group("/v1")

	{
		v1.GET("request", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Many Requests"})
		})

	}

	return r
}
