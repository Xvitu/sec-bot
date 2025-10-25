package repository

import (
	"xvitu/sec-bot/domain"
	"xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/infra/persistence"
)

type MessageRepository struct{}

func (r *MessageRepository) GetByStepAndMessageId(step domain.Step, messageId string) entity.Message {
	return entity.Message{
		Id:   messageId,
		Text: persistence.Messages[step][messageId],
	}
}
