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

func sendMeme(bot *tgbotapi.BotAPI, channel string, workTime WorkTime) {
	if workTime.IsWorkTime == workTime.CheckTimePeriod(time.Now()) {
		return
	}
	meme := getMeme()
	message := tgbotapi.NewPhotoToChannel(channel, tgbotapi.FileURL(meme.URL))
	bot.Send(message)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	start, _ := time.Parse("15:00", "22:00")
	end, _ := time.Parse("15:00", "09:00")
	workTime := WorkTime{false, start, end}

	time.Parse("15:00", "20")

	botToken, _ := os.LookupEnv("BOT_TOKEN")
	channelId, _ := os.LookupEnv("CHANNEL")

	bot, err := configBot(botToken)
	if err != nil {
		log.Panicln(err)
	}

	go startBot(bot)

	sendMeme(bot, channelId, workTime)

	ticker := time.NewTicker(60 * time.Minute)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				sendMeme(bot, channelId, workTime)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	for {
	}
}
