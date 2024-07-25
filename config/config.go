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

func LoadBotConfig() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	botToken := os.Getenv("TELEGRAM_TOKEN")
	if botToken == "" {
		log.Fatalf("TELEGRAM_TOKEN not set in .env file")
	}
	return botToken
}
