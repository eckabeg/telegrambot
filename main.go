package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

const apiKey = "609285927:AAFDVQsbNPMFlo2rMNUJzkNkYrZcpIDS6H8"

func main() {
	port := os.Getenv("PORT")
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://vast-mountain-33032.herokuapp.com/"+bot.Token, nil))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe("0.0.0.0:"+port, nil)
	MsgInvover := MessageInvoker{}
	InlineInvoker := InlineInvoker{}

	for update := range updates {
		log.Printf("%+v\n", update)
		if update.Message != nil {
			MsgInvover.HandleMessage(bot, *update.Message)
		} else if update.InlineQuery != nil {
			InlineInvoker.HandleMessage(bot, *update.InlineQuery)
		}

	}
}
