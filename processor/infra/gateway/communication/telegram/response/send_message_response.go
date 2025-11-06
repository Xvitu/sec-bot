package response

import (
	"github.com/xvitu/sec-bot/processor/shared/types"
)

type Chat struct {
	ID   types.String `json:"id"`
	Type types.String `json:"type"`
}

type Message struct {
	MessageID types.String `json:"message_id"`
	Chat      Chat         `json:"chat"`
	Text      types.String `json:"text"`
}

type SendMessageResponse struct {
	Result Message `json:"result"`
}
