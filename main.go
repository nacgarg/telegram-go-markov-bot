package main 

import (
    "log"
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "flag"
    "strings"
    "fmt"
)

var dataDict = make(map[string]map[string]int)
var filePath string

func main() {
	// Read flags

	flag.StringVar(&filePath, "file", "data.mrkv", "Path to markov dataset file")

	var botToken string
	flag.StringVar(&botToken, "token", "", "Telegram token for bot")

	var importFile string
	flag.StringVar(&importFile, "import", "", "Import raw text logs")

	flag.Parse()


	load_dataset(filePath)

	if importFile != "" {
		import_file(importFile)
	}

	// Init bot


    bot, err := tgbotapi.NewBotAPI(botToken)
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true

    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates, err := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil || update.Message.Text == "" {
            continue
        }
        if strings.Fields(update.Message.Text)[0] == "/test" {
        	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "yo")
       		msg.ReplyToMessageID = update.Message.MessageID
       		bot.Send(msg)
        }

         if strings.Fields(update.Message.Text)[0] == "/markov" {
         	fmt.Println(update.Message.Text[8:])
        	msg := tgbotapi.NewMessage(update.Message.Chat.ID, generate_response(update.Message.Text[8:]))
       		msg.ReplyToMessageID = update.Message.MessageID
       		bot.Send(msg)
        }

        train(update.Message.Text)

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)


    }
}

