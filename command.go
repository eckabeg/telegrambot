package main

import (
	"gopkg.in/telegram-bot-api.v4"
)

type Command struct {
	bot    *tgbotapi.BotAPI
	m      tgbotapi.Message
	params []string
}
