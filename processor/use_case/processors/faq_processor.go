package processors

import (
	"context"
	"fmt"
	"xvitu/sec-bot/domain"
	domainEntity "xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/entypoint/dto"
	"xvitu/sec-bot/infra/gateway/communication"
	"xvitu/sec-bot/infra/persistence/repository"
)

type FaqProcessor struct {
	chatRepository       repository.ChatRepositoryInterface
	communicationGateway communication.CommunicationGateway
	messageRepository    repository.MessageRepositoryInterface
}

const backMessageId = "11"

func NewFaqProcessor(
	chatRepository repository.ChatRepositoryInterface,
	communicationGateway communication.CommunicationGateway,
	messageRepository repository.MessageRepositoryInterface,
) *FaqProcessor {
	return &FaqProcessor{
		chatRepository:       chatRepository,
		communicationGateway: communicationGateway,
		messageRepository:    messageRepository,
	}
}

func (p *FaqProcessor) Execute(chatUpdate dto.Chat, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	if chatUpdate.Message == backMessageId {
		return p.handleMessage(domain.MainMenu, "greetings", chat)
	}

	updatedChat, chatError := p.handleMessage(domain.Faq, "faq_"+chatUpdate.Message, chat)
	if chatError != nil {
		return p.handleError("invalid_option", chat)
	}

	return updatedChat, nil
}

func (p *FaqProcessor) handleError(messageId string, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	replyMessage := p.messageRepository.GetByStepAndMessageId(domain.Error, messageId)
	_, sendMessageError := p.communicationGateway.SendMessage(chat.ExternalId, replyMessage.Text)
	if sendMessageError != nil {
		return nil, fmt.Errorf("error while sending message: %s", sendMessageError)
	}

	return chat, nil
}

func (p *FaqProcessor) handleMessage(step domain.Step, messageId string, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	replyMessage := p.messageRepository.GetByStepAndMessageId(step, messageId)
	if replyMessage == nil {
		return nil, fmt.Errorf("no Message Found")
	}

	_, sendMessageError := p.communicationGateway.SendMessage(chat.ExternalId, replyMessage.Text)
	if sendMessageError != nil {
		return nil, fmt.Errorf("error while sending message: %s", sendMessageError)
	}

	chat.Step = step
	chat.LastMessageID = messageId
	saveError := p.chatRepository.Save(context.Background(), *chat)
	if saveError != nil {
		return nil, fmt.Errorf("error while saving chat: %s", saveError)
	}
	return chat, nil
}
