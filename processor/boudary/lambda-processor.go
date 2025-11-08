package boudary

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/xvitu/sec-bot/processor/application/use_case"
	"github.com/xvitu/sec-bot/processor/entypoint/dto"
	"github.com/xvitu/sec-bot/processor/shared/json"
)

type LambdaProcessor struct {
	chatUpdateHandler *use_case.ChatUpdateHandler
}

func NewLambdaProcessor(handler *use_case.ChatUpdateHandler) *LambdaProcessor {
	return &LambdaProcessor{
		chatUpdateHandler: handler,
	}
}

func (l *LambdaProcessor) Handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		chatDto, _ := json.ToStruct[dto.Chat]([]byte(message.Body))
		l.chatUpdateHandler.Run(chatDto)
	}

	return nil
}
