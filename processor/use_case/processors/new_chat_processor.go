package processors

import (
	"context"
	"time"
	"xvitu/sec-bot/domain"
	domainEntity "xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/entypoint/dto"
	"xvitu/sec-bot/infra/gateway/communication"
	"xvitu/sec-bot/infra/persistence/repository"

	"github.com/google/uuid"
)

type NewChatProcessor struct {
	chatRepository       repository.ChatRepositoryInterface
	communicationGateway communication.CommunicationGateway
	messageRepository    repository.MessageRepositoryInterface
}

func CreateNewChatProcessor(
	chatRepository repository.ChatRepositoryInterface,
	communicationGateway communication.CommunicationGateway,
	messageRepository repository.MessageRepositoryInterface,
) *NewChatProcessor {
	return &NewChatProcessor{
		chatRepository:       chatRepository,
		communicationGateway: communicationGateway,
		messageRepository:    messageRepository,
	}
}

func (p *NewChatProcessor) Execute(chatUpdate dto.Chat, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	message := p.messageRepository.GetByStepAndMessageId(domain.MainMenu, "greetings")
	_, sendMessageError := p.communicationGateway.SendMessage(chatUpdate.ExternalId, message.Text)

	if sendMessageError != nil {
		return nil, sendMessageError
	}

	domainChat := domainEntity.Chat{
		Origin:        chatUpdate.Origin,
		ExternalId:    chatUpdate.ExternalId,
		UserId:        chatUpdate.ExternalUserId,
		CreatedAt:     time.Now().String(),
		UpdatedAt:     time.Now().String(),
		Step:          domain.MainMenu,
		LastMessageID: message.Id,
		Id:            uuid.New().String(),
	}
	createErr := p.chatRepository.Save(context.Background(), domainChat)

	if createErr != nil {
		return nil, createErr
	}

	return &domainChat, nil
}
