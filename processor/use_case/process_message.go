package use_case

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

type ChatUpdateProcessor struct {
	chatRepository    *repository.ChatRepository
	telegramGateway   *telegram.Gateway
	messageRepository *repository.MessageRepository
}

func NewChatUpdateProcessor(
	chatRepository *repository.ChatRepository,
	telegramGateway *telegram.Gateway,
	messageRepository *repository.MessageRepository,
) *ChatUpdateProcessor {
	return &ChatUpdateProcessor{
		chatRepository:    chatRepository,
		telegramGateway:   telegramGateway,
		messageRepository: messageRepository,
	}
}

func (u *ChatUpdateProcessor) Run(chatUpdate dto.Chat) (*domainEntity.Chat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	persistedChat, chatRepoError := u.chatRepository.FindByID(ctx, chatUpdate.ExternalUserId)

	if chatRepoError != nil {
		return nil, chatRepoError
	}

	if persistedChat == nil {
		message := u.messageRepository.GetByStepAndMessageId(domain.Start, "greetings")
		_, sendMessageError := u.telegramGateway.SendMessage(chatUpdate.ExternalId, message.Text)

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
		createErr := u.chatRepository.Save(context.Background(), domainChat)

		if createErr != nil {
			return nil, createErr
		}

		return &domainChat, nil
	}

	return nil, nil
}
