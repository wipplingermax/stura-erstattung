package config

import (
	"fmt"
	"io"
	"log"
	"os"
)

func InitLogToFile() {
	// If the file doesn't exist, create it or append to the file
	var path string = "/var/log"
	var filename string = "stura-erstattung.log"

	if os.Getenv("LOG_PATH") != "" {
		path = os.Getenv("LOG_PATH")
	}

	if os.Getenv("LOG_FILE") != "" {
		filename = os.Getenv("LOG_FILE")
	}

	file, err := os.OpenFile(
		fmt.Sprintf("%s/%s", path, filename),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)

	if err != nil {
		log.Printf("Failed to Log into File: %s", err)
		return
	}

	log.SetOutput(io.MultiWriter(os.Stderr, file))
}
