package communication

import "github.com/xvitu/sec-bot/processor/infra/gateway/communication/response"

type CommunicationGateway interface {
	SendMessage(chatID string, text string) (*response.SendMessageResponse, error)
}
