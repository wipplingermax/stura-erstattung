package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

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

	return r
}
