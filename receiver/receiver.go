package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/xvitu/sec-bot/receiver/boundary"
	"github.com/xvitu/sec-bot/receiver/boundary/middleware"
	"github.com/xvitu/sec-bot/receiver/infra/sqs"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	client := (sqs.SqsClient{})
	sqsClient := client.Create()
	webhookController := boundary.NewWebhookController(sqsClient)

	http.Handle(
		"/telegram/webhook/"+os.Getenv("TELEGRAM_WEBHOOK_TOKEN"),
		middleware.ValidateTelegramIP(http.HandlerFunc(webhookController.HandleRequest)),
	)

	fmt.Println("Running bot on port 8080")
	http.ListenAndServe(":8080", nil)

}
