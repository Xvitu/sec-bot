package request

import "encoding/json"

type SendMessageRequest struct {
	ChatID    int64  `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}

func NewSendMessageRequest(chatId int64, text string) SendMessageRequest {
	return SendMessageRequest{
		ChatID:    chatId,
		Text:      text,
		ParseMode: "Markdown",
	}
}

func (r SendMessageRequest) Endpoint() string {
	return "/sendMessage"
}

func (r SendMessageRequest) Body() ([]byte, error) {
	return json.Marshal(r)
}
