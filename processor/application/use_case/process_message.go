package use_case

import (
	"context"
	"time"
	"xvitu/sec-bot/application/use_case/processors"
	"xvitu/sec-bot/domain"
	domainEntity "xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/entypoint/dto"
	"xvitu/sec-bot/infra/persistence/repository"
)

type ChatUpdateHandler struct {
	messageProcessors map[domain.Step]processors.MessageProcessor
	chatRepository    repository.ChatRepositoryInterface
}

func NewChatUpdateHandler(messageProcessors map[domain.Step]processors.MessageProcessor, repositoryInterface repository.ChatRepositoryInterface) *ChatUpdateHandler {
	return &ChatUpdateHandler{
		messageProcessors: messageProcessors,
		chatRepository:    repositoryInterface,
	}
}

func (u *ChatUpdateHandler) Run(chatUpdate dto.Chat) (*domainEntity.Chat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	persistedChat, chatRepoError := u.chatRepository.FindByExternalId(ctx, chatUpdate.ExternalUserId)

	if chatRepoError != nil {
		return nil, chatRepoError
	}

	if persistedChat == nil {
		return u.messageProcessors[domain.Start].Execute(chatUpdate, persistedChat)
	}

	processor := u.messageProcessors[persistedChat.Step]
	return processor.Execute(chatUpdate, persistedChat)
}

// todo - consts para id das mensgens?
