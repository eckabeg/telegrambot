package main

import (
	"reflect"
	"strings"

	"gopkg.in/telegram-bot-api.v4"
)

type MessageInvoker struct {
}

/*
Handles Incoming Messages and dispatches them to the available functions
(Happens with reflectionmagic.)
*/
func (m MessageInvoker) HandleMessage(b *tgbotapi.BotAPI, msg tgbotapi.Message) {
	if msg.IsCommand() {
		var command string
		var params []string
		if strings.Contains(msg.Text, " ") {
			command = strings.Title(strings.ToLower(msg.Text[1:strings.Index(msg.Text, " ")]))
		} else {
			command = strings.Title(strings.ToLower(msg.Text[1:]))
		}
		if strings.Contains(msg.Text, " ") {
			params = strings.SplitN(msg.Text[strings.Index(msg.Text, " ")+1:], " ", -1)
		}
		c := Command{b, msg, params}
		f := reflect.ValueOf(&c).MethodByName(command)
		if f.IsValid() {
			f.Call([]reflect.Value{})
		} else {
			c.Help()
		}

	} else {
		println(msg.Text)
	}
}

// Help Sends a Message with all available Commands
func (c Command) Help() {
	rc := reflect.TypeOf(&c)
	s := "Ich biete dir folgende Funktionen an: \n"
	for i := 0; i < rc.NumMethod(); i++ {
		s += "/" + rc.Method(i).Name
		s += "\n"
	}
	msg := tgbotapi.NewMessage(c.m.Chat.ID, s)
	c.bot.Send(msg)
}

func getAllCommands() []string {
	var c Command
	var returnCommands []string
	rc := reflect.TypeOf(&c)
	for i := 0; i < rc.NumMethod(); i++ {
		returnCommands = append(returnCommands, rc.Method(i).Name)
	}
	return returnCommands
}
