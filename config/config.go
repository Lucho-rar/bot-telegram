package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() (string, string) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalf("API_KEY not set in .env file")
	}

	spreadsheetId := os.Getenv("SPREADSHEET_ID")
	if spreadsheetId == "" {
		log.Fatalf("SPREADSHEET_ID not set in .env file")
	}
	return apiKey, spreadsheetId
}
