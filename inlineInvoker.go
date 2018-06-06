package main

import "gopkg.in/telegram-bot-api.v4"
import "strconv"

type InlineInvoker struct {
}

/*
Handles Incoming Messages and dispatches them to the available functions
(Happens with reflectionmagic.)
*/
func (m InlineInvoker) HandleMessage(b *tgbotapi.BotAPI, msg tgbotapi.InlineQuery) {
	println(msg.Offset)
	println(msg.Query)

	helparticles := buildCommandArticles(msg.ID)

	res := make([]interface{}, len(helparticles))
	for i, v := range helparticles {
		res[i] = v
	}

	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: msg.ID,
		IsPersonal:    true,
		CacheTime:     0,
		Results:       res,
	}

	b.AnswerInlineQuery(inlineConf)
}

func buildCommandArticles(msgId string) []tgbotapi.InlineQueryResultArticle {
	var resultArticles []tgbotapi.InlineQueryResultArticle

	for i, c := range getAllCommands() {
		article := tgbotapi.NewInlineQueryResultArticle(strconv.Itoa(i), "Commands", c)
		article.Description = c
		resultArticles = append(resultArticles, article)
	}
	return resultArticles

}

// Help Sends a Message with all available Commands
