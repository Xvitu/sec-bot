package telegram

import (
	client "github.com/xvitu/sec-bot/processor/infra/client/telegram"
	"github.com/xvitu/sec-bot/processor/infra/gateway/communication/response"
	"github.com/xvitu/sec-bot/processor/infra/gateway/communication/telegram/request"
	telegramResponse "github.com/xvitu/sec-bot/processor/infra/gateway/communication/telegram/response"
	"github.com/xvitu/sec-bot/processor/shared/json"
	"strconv"
)

type Gateway struct {
	client *client.Client
}

func NewGateway(client *client.Client) *Gateway {
	return &Gateway{client: client}
}

func (g *Gateway) SendMessage(chatID string, text string) (*response.SendMessageResponse, error) {
	intChatId, _ := strconv.ParseInt(chatID, 10, 64)
	sendMessageRequest := request.NewSendMessageRequest(intChatId, text)
	responseBytes, postErr := g.client.Post(sendMessageRequest)

	if postErr != nil {
		return nil, postErr
	}

	resultStruct, marshalErr := json.ToStruct[telegramResponse.SendMessageResponse](responseBytes)

	if marshalErr != nil {
		return nil, marshalErr
	}

	return &response.SendMessageResponse{
		MessageId: resultStruct.Result.MessageID.Get(),
		ChatId:    resultStruct.Result.Chat.ID.Get(),
		Text:      resultStruct.Result.Text.Get(),
	}, nil
}
