package config

import "os"

func Init() {

	// enable Development environment
	if os.Getenv("DEVELOPMENT") == "true" {
		InitEnv()
	}

	// enable additional logging to file
	if os.Getenv("LOG_TO_FILE") == "true" {
		InitLogToFile()
	}
}
