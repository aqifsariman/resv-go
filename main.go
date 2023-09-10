package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func init() {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	// Build path to .env in the current directory
	envPath := filepath.Join(cwd, ".env")

	// Load the .env file
	if err := godotenv.Load(envPath); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	fmt.Println("Running RSV.")
}
