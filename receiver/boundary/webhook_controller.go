package boundary

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/xvitu/sec-bot/receiver/boundary/request"
	"github.com/xvitu/sec-bot/receiver/domain"
	sqs2 "github.com/xvitu/sec-bot/receiver/infra/sqs"
)

type WebHookController struct {
	sqsClient *sqs.Client
}

func NewWebhookController(
	sqsClient *sqs.Client,
) *WebHookController {
	return &WebHookController{
		sqsClient: sqsClient,
	}
}

func (c *WebHookController) HandleRequest(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	token := req.PathParameters["token"]

	if token != os.Getenv("TELEGRAM_WEBHOOK_TOKEN") {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 403,
			Body:       `{"error":"invalid token"}`,
		}, nil
	}

	body := req.Body

	var update request.ChatUpdateRequest
	if err := json.Unmarshal([]byte(body), &update); err != nil {
		fmt.Println("Erro ao fazer unmarshal:", err)
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       `{"error":"invalid json"}`,
		}, nil
	}

	queue := os.Getenv("UPDATE_CHAT_QUEUE")

	chatDto := sqs2.Chat{
		ExternalId:     strconv.FormatInt(update.Message.Chat.Id, 10),
		ExternalUserId: strconv.FormatInt(update.Message.From.Id, 10),
		SentAt:         time.Now().String(),
		Message:        update.Message.Text,
		Origin:         domain.TelegramOrigin,
	}

	jsonMessage, _ := json.Marshal(chatDto)

	out, err := c.sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String(string(jsonMessage)),
		QueueUrl:    aws.String(queue),
	})

	fmt.Println("SendMessage OUT:", out, "ERR:", err)

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jsonMessage),
	}, nil
}
