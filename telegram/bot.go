package telegram

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Lucho-rar/bot-telegram/data"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/api/sheets/v4"
)

func StartBot(t string, srv *sheets.Service, spreadsheetId string) {
	bot, err := tgbotapi.NewBotAPI(t)
	if err != nil {
		log.Fatalf("Failed to create Telegram bot: %v", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.Text = "Hi! I'm irregular verbs Bot! Come on practice with me! 💯"
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		msg.Text = "Te voy a mandar un verbo!"
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(224) + 1

		readRange := fmt.Sprintf("bot-telegram!A%d:A%d", num, num)
		//readRange := "bot-telegram!A2:A224"
		datos, err := data.ReadData(srv, spreadsheetId, readRange)
		if err != nil {
			log.Fatalf("Unable to retrieve data from sheet: %v", err)
		}

		verb := data.ReturnData(datos)
		//hoja.PrintData(data)
		msg.Text = "Infinitivo: " + verb
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

		readRange = fmt.Sprintf("bot-telegram!B%d:B%d", num, num)
		//readRange := "bot-telegram!A2:A224"
		datos, err = data.ReadData(srv, spreadsheetId, readRange)
		if err != nil {
			log.Fatalf("Unable to retrieve data from sheet: %v", err)
		}

		msg.Text = "15 segundos!⏰"
		bot.Send(msg)
		wait := time.Duration(15) * time.Second
		time.Sleep(wait)

		verb = data.ReturnData(datos)
		//hoja.PrintData(data)
		msg.Text = "Pasado simple: " + verb
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

		readRange = fmt.Sprintf("bot-telegram!C%d:C%d", num, num)
		//readRange := "bot-telegram!A2:A224"
		datos, err = data.ReadData(srv, spreadsheetId, readRange)
		if err != nil {
			log.Fatalf("Unable to retrieve data from sheet: %v", err)
		}

		verb = data.ReturnData(datos)
		//hoja.PrintData(data)
		msg.Text = "Participio pasado :" + verb
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}

func HandlerBot() {

}
