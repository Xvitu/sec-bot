package communication

import "xvitu/sec-bot/infra/gateway/communication/response"

type CommunicationGateway interface {
	SendMessage(chatID string, text string) (*response.SendMessageResponse, error)
}
