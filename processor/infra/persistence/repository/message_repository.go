package repository

import (
	"xvitu/sec-bot/domain"
	"xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/infra/persistence"
)

type MessageRepositoryInterface interface {
	GetByStepAndMessageId(step domain.Step, messageId string) *entity.Message
}

type MessageRepository struct{}

func (r *MessageRepository) GetByStepAndMessageId(step domain.Step, messageId string) *entity.Message {
	messageStep := persistence.Messages[step]

	if messageStep != nil {
		message := messageStep[messageId]

		if message != "" {
			return &entity.Message{
				Id:   messageId,
				Text: message,
			}
		}
	}

	return nil
}
