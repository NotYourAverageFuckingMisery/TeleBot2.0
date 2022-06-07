package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"Telebot2.0/modules"
)

// this uses telegram bot API they have an amasing documentation

func GetUpdates(bot_url string, offset int) ([]modules.Update, error) {
	resp, err := http.Get(bot_url + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse modules.RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

func BotPost(botMsg string, bot_url string, update modules.Update) error {
	var botMessage modules.BotMessage
	botMessage.Text = botMsg
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.ReplyMsgId = update.Message.MsgId

	buf, err := json.Marshal(botMessage)
	if err != nil {
		panic(err.Error)
	}

	_, err = http.Post(bot_url+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		panic(err.Error())
	}
	return nil
}
