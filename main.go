package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"os"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		msg = updateMsg(msg)

		bot.Send(msg)
	}
}

func updateMsg(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig {
	//var addons = []string{", но едя пидор", ", но едя хуй"}

	msg.Text += ", но едя пидор"

	return msg
}