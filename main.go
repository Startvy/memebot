package main

import (
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func getMeme() Responce {
	var api MemeApi
	api = Api{BaseUrl: "https://meme-api.herokuapp.com"}
	meme := api.GetMeme()

	return meme
}

func sendMeme(bot *tgbotapi.BotAPI, channel string) {
	meme := getMeme()
	// message := tgbotapi.NewPhoto(MASTER_CHAT_ID, tgbotapi.FileURL(meme.URL))
	message := tgbotapi.NewPhotoToChannel("@memetikovo", tgbotapi.FileURL(meme.URL))
	bot.Send(message)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	botToken, _ := os.LookupEnv("BOT_TOKEN")
	channelId, _ := os.LookupEnv("CHANNEL")

	bot, err := configBot(botToken)
	if err != nil {
		log.Panicln(err)
	}

	go startBot(bot)

	sendMeme(bot, channelId)

	ticker := time.NewTicker(30 * time.Minute)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				sendMeme(bot, channelId)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	for {
	}
}
