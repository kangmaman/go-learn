package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	username, exists := os.LookupEnv("USSI_USERNAME")
	if exists {
		fmt.Printf("USSI_USERNAME: %s\n", username)
	}

	apiKey, exists := os.LookupEnv("USSI_API_KEY")
	if exists {
		fmt.Printf("USSI_API_KEY: %s\n", apiKey)
	}
}

func init() {
	//loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
