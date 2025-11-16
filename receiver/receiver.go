package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/xvitu/sec-bot/receiver/boundary"
	"github.com/xvitu/sec-bot/receiver/infra/sqs"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	client := sqs.SqsClient{}
	sqsClient := client.Create(nil)
	webhookController := boundary.NewWebhookController(sqsClient)

	lambda.Start(webhookController.HandleRequest)
}
