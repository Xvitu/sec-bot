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

func (c *WebHookController) HandleRequest(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	token := req.PathParameters["token"]

	if token != os.Getenv("TELEGRAM_WEBHOOK_TOKEN") {
		return events.APIGatewayProxyResponse{
			StatusCode: 403,
			Body:       `{"error":"invalid token"}`,
		}, nil
	}

	var update request.ChatUpdateRequest
	if err := json.Unmarshal([]byte(req.Body), &update); err != nil {
		fmt.Println("Erro ao fazer unmarshal:", err)
		return events.APIGatewayProxyResponse{
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

	jsonMessage, err := json.Marshal(chatDto)
	if err != nil {
		fmt.Println("Erro ao serializar mensagem SQS:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       `{"error":"internal error"}`,
		}, nil
	}

	out, err := c.sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String(string(jsonMessage)),
		QueueUrl:    aws.String(queue),
	})

	fmt.Println("SendMessage OUT:", out, "ERR:", err)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       `{"error":"failed to send message to SQS"}`,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jsonMessage),
	}, nil
}
