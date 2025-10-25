package telegram

import (
	"xvitu/sec-bot/infra/telegram/request"
	"xvitu/sec-bot/infra/telegram/response"
	"xvitu/sec-bot/shared/json"
)

type Gateway struct {
	Client *Client
}

func NewGateway(client *Client) *Gateway {
	return &Gateway{Client: client}
}

func (g *Gateway) SendMessage(chatID int64, text string) (*response.SendMessageResponse, error) {
	sendMessageRequest := request.NewSendMessageRequest(chatID, text)
	responseBytes, postErr := g.Client.Post(sendMessageRequest)

	if postErr != nil {
		return nil, postErr
	}

	resultStruct, marshalErr := json.ToStruct[response.SendMessageResponse](responseBytes)

	if marshalErr != nil {
		return nil, marshalErr
	}

	return resultStruct, nil
}
