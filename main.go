package main

import (
	"strings"

	"Telebot2.0/controllers"
)

const bot_token = "YOUR_BOT_TOKEN"
const bot_api = "https://api.telegram.org/bot"
const bot_url = bot_api + bot_token

func main() {

	offset := 0
	marker := false

	for {
		updates, err := controllers.GetUpdates(bot_url, offset)
		if err != nil {
			panic(err.Error())
		}

		for _, update := range updates {

			//controllers.BotPost(update.Message.Text, bot_url, update) this is an echo bot function. turn on for testing/debug

			if marker {
				city := update.Message.Text
				lat, lon := controllers.GetGeocode(city)
				weather := controllers.GetWeather(lat, lon)
				controllers.BotPost(weather, bot_url, update)
				offset = update.UpdateId + 1
				marker = false
			}

			if strings.ToLower(update.Message.Text) == "/weather" {
				controllers.BotPost("Enter your city:", bot_url, update)
				marker = true
				offset = update.UpdateId + 1
			}
		}

	}
}
