package main

import (
	"log"
	//	"github.c/config" // Replace "your-package-path" with the actual package path
	"github.com/Lucho-rar/bot-telegram/config"
	"github.com/Lucho-rar/bot-telegram/sheets"
)

func main() {
	// config
	apiKey, spreadsheetId := config.LoadConfig()

	// init
	readRange := "bot-telegram!A:D"

	srv, err := sheets.NewService(apiKey)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Lread
	data, err := sheets.ReadData(srv, spreadsheetId, readRange)
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	// print
	sheets.PrintData(data)
}
