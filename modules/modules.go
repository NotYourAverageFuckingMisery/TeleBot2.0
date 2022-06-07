package modules

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Text  string `json:"text"`
	Chat  Chat   `json:"chat"`
	MsgId int    `json:"message_id"`
}

type Chat struct {
	ChatId int `json:"id"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatId     int    `json:"chat_id"`
	Text       string `json:"text"`
	ReplyMsgId int    `json:"reply_to_message_id"`
}

// I love the room movie it's so bad and so good
