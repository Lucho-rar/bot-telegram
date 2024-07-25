package main

import (
	"log"

	"github.com/Lucho-rar/bot-telegram/config"
	"github.com/Lucho-rar/bot-telegram/data"
	"github.com/Lucho-rar/bot-telegram/telegram"
)

func main() {
	// config
	apiKey, spreadsheetId := config.LoadConfig()

	srv, err := data.NewService(apiKey)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	botToken := config.LoadBotConfig()
	telegram.StartBot(botToken, srv, spreadsheetId)
}
