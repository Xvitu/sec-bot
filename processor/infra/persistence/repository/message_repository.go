package repository

import (
	"xvitu/sec-bot/domain"
	"xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/infra/persistence"
)

func GetByStepAndMessageId(step domain.Step, messageId string) entity.Message {
	return entity.Message{
		Id:   messageId,
		Text: persistence.Messages[step][messageId],
	}
}
