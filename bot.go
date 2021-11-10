package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func startBot(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		message := fmt.Sprintf("ID: %d \n Name: %s", update.Message.Chat.ID, update.Message.From.UserName)
		log.Printf("ID: %d \nName: %s", update.Message.Chat.ID, update.Message.From.UserName)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func configBot(token string) (bot *tgbotapi.BotAPI, err error) {
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return
}
