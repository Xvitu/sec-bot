package boundary

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

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

func (c *WebHookController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	var update request.ChatUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "JSON inválido", http.StatusUnprocessableEntity)
		return
	}

	queue := os.Getenv("UPDATE_CHAT_QUEUE")
	ctx := context.TODO()

	chatDto := sqs2.Chat{
		ExternalId:     strconv.FormatInt(update.Message.Chat.Id, 10),
		ExternalUserId: strconv.FormatInt(update.Message.From.Id, 10),
		SentAt:         time.Now().String(),
		Message:        update.Message.Text,
		Origin:         domain.TelegramOrigin,
	}

	jsonMessage, _ := json.Marshal(chatDto)

	c.sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String(string(jsonMessage)),
		QueueUrl:    aws.String(queue),
	})

	fmt.Printf("Mensagem recebida: %s (Chat ID: %d)\n", update.Message.Text, update.Message.Chat.Id)

	w.WriteHeader(http.StatusOK)
}
