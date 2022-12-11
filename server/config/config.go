package config

import (
	"fmt"
	"os"
)

func Init() {

	// enable Development environment
	if _, err := os.Stat(".env"); !os.IsNotExist(err) {
		InitEnvironment()
		fmt.Println("loaded local .env file next to executable")
	}

	// enable additional logging to file
	if os.Getenv("LOG_TO_FILE") == "true" {
		InitLogging()
	}
}
