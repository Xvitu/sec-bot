package telegram

import (
	"strconv"
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

func (g *Gateway) SendMessage(chatID string, text string) (*response.SendMessageResponse, error) {
	intChatId, _ := strconv.ParseInt(chatID, 10, 64)
	sendMessageRequest := request.NewSendMessageRequest(intChatId, text)
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
