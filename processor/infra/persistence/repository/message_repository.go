package repository

import (
	"github.com/xvitu/sec-bot/processor/domain"
	"github.com/xvitu/sec-bot/processor/domain/entity"
	"github.com/xvitu/sec-bot/processor/infra/persistence"
	"slices"
)

type MessageRepositoryInterface interface {
	GetByStepAndMessageId(step domain.Step, messageId string) *entity.Message
	FindAllByStepExcludingIds(step domain.Step, messagesIds []string) []*entity.Message
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

func (r *MessageRepository) FindAllByStepExcludingIds(step domain.Step, messagesIds []string) []*entity.Message {
	allStepMessages := persistence.Messages[step]

	var filteredMessages []*entity.Message
	for messageId, message := range allStepMessages {
		if !slices.Contains(messagesIds, messageId) {
			filteredMessages = append(filteredMessages, &entity.Message{
				Id:   messageId,
				Text: message,
			})
		}
	}

	return filteredMessages
}
