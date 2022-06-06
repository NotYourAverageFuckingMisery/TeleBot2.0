package main

import "Telebot2.0/controllers"

const bot_token = "5434054221:AAEq7SL23w3uf5g2irIFMEX1tQ4QCO3c4jo"
const bot_api = "https://api.telegram.org/bot"
const bot_url = bot_api + bot_token

func main() {

	offset := 0

	for {
		updates, err := controllers.GetUpdates(bot_url, offset)
		if err != nil {
			panic(err.Error())
		}

		for _, update := range updates {
			controllers.BotPost(update.Message.Text, bot_url, update)
			offset = update.UpdateId + 1
		}

	}
}
