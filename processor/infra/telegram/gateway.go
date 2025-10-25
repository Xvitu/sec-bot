package telegram

import (
	"xvitu/sec-bot/infra/telegram/request"
	"xvitu/sec-bot/infra/telegram/response"
	"xvitu/sec-bot/shared/json"
)

type Gateway struct {
	client *Client
}

func NewGateway(client *Client) *Gateway {
	return &Gateway{client: client}
}

func (g *Gateway) SendMessage(chatID int64, text string) (*response.SendMessageResponse, error) {
	sendMessageRequest := request.NewSendMessageRequest(chatID, text)
	responseBytes, postErr := g.client.Post(sendMessageRequest)

	if postErr != nil {
		return nil, postErr
	}

	resultStruct, marshalErr := json.ToStruct[response.SendMessageResponse](responseBytes)

	if marshalErr != nil {
		return nil, marshalErr
	}

	return resultStruct, nil
}
