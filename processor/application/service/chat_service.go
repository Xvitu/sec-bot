package service

import (
	"context"
	"fmt"

	"github.com/xvitu/sec-bot/processor/domain"
	domainEntity "github.com/xvitu/sec-bot/processor/domain/entity"
	"github.com/xvitu/sec-bot/processor/infra/gateway/communication"
	"github.com/xvitu/sec-bot/processor/infra/persistence/repository"
)

type ChatService struct {
	chatRepository       repository.ChatRepositoryInterface
	communicationGateway communication.CommunicationGateway
	messageRepository    repository.MessageRepositoryInterface
}

func NewChatService(
	chatRepository repository.ChatRepositoryInterface,
	communicationGateway communication.CommunicationGateway,
	messageRepository repository.MessageRepositoryInterface,
) *ChatService {
	return &ChatService{
		chatRepository:       chatRepository,
		communicationGateway: communicationGateway,
		messageRepository:    messageRepository,
	}
}

func (s *ChatService) HandleError(messageId string, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	replyMessage := s.messageRepository.GetByStepAndMessageId(domain.Error, messageId)
	_, sendMessageError := s.communicationGateway.SendMessage(chat.ExternalId, replyMessage.Text)
	if sendMessageError != nil {
		return nil, fmt.Errorf("error while sending message: %s", sendMessageError)
	}

	return chat, nil
}

func (s *ChatService) HandleReplyMessages(step domain.Step, messageIds []string, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	for _, messageId := range messageIds {
		replyMessage := s.messageRepository.GetByStepAndMessageId(step, messageId)
		if replyMessage == nil {
			return nil, fmt.Errorf("no Message Found")
		}

		_, sendMessageError := s.communicationGateway.SendMessage(chat.ExternalId, replyMessage.Text)
		if sendMessageError != nil {
			return nil, fmt.Errorf("error while sending message: %s", sendMessageError)
		}
	}

	chat.UpdateWithRepledMessage(messageIds[0], step)
	saveError := s.chatRepository.Save(context.Background(), *chat)
	if saveError != nil {
		return nil, fmt.Errorf("error while saving chat: %s", saveError)
	}
	return chat, nil
}
