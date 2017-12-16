package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"net/http"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("69024058:AAHv4XRYX-pyr_-jmnv8PBcrytfrg1OWORo")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, "cert.pem"))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}