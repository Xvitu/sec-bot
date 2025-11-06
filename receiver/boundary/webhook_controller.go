package boundary

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/xvitu/sec-bot/receiver/boundary/request"
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
	c.sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String(`{"update_id": 123456, "message": {"text": "Olá da fila!"}}`),
		QueueUrl:    aws.String(queue),
	})

	fmt.Printf("Mensagem recebida: %s (Chat ID: %d)\n", update.Message.Text, update.Message.Chat.ID)

	w.WriteHeader(http.StatusOK)
}
