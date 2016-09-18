package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

var CommandMap = map[string]func(*tgbotapi.BotAPI, tgbotapi.Update){
	"markov": handleMarkov,
	"test":   handleTest,
}

func handleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	fn, ok := CommandMap[strings.ToLower(update.Message.Command())]

	if !ok {
		return
	}

	fn(bot, update)
}

func handleTest(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "yo")

	msg.ReplyToMessageID = update.Message.MessageID

	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func handleMarkov(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, generateMarkovResponse(update.Message.CommandArguments()))

	msg.ReplyToMessageID = update.Message.MessageID

	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}