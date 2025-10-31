package processors

import (
	"context"
	"time"
	"xvitu/sec-bot/domain"
	domainEntity "xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/entypoint/dto"
	"xvitu/sec-bot/infra/persistence/repository"
	"xvitu/sec-bot/infra/telegram"

	"github.com/google/uuid"
)

type NewChatProcessor struct {
	chatRepository    repository.ChatRepositoryInterface
	telegramGateway   *telegram.Gateway
	messageRepository repository.MessageRepositoryInterface
}

func CreateNewChatProcessor(
	chatRepository repository.ChatRepositoryInterface,
	telegramGateway *telegram.Gateway,
	messageRepository repository.MessageRepositoryInterface,
) *NewChatProcessor {
	return &NewChatProcessor{
		chatRepository:    chatRepository,
		telegramGateway:   telegramGateway,
		messageRepository: messageRepository,
	}
}

func (p *NewChatProcessor) Execute(chatUpdate dto.Chat, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	message := p.messageRepository.GetByStepAndMessageId(domain.Start, "greetings")
	_, sendMessageError := p.telegramGateway.SendMessage(chatUpdate.ExternalId, message.Text)

	if sendMessageError != nil {
		return nil, sendMessageError
	}

	domainChat := domainEntity.Chat{
		Origin:        chatUpdate.Origin,
		ExternalId:    chatUpdate.ExternalId,
		UserId:        chatUpdate.ExternalUserId,
		CreatedAt:     time.Now().String(),
		UpdatedAt:     time.Now().String(),
		Step:          domain.Start,
		LastMessageID: message.Id,
		Id:            uuid.New().String(),
	}
	createErr := p.chatRepository.Save(context.Background(), domainChat)

	if createErr != nil {
		return nil, createErr
	}

	return &domainChat, nil

}
