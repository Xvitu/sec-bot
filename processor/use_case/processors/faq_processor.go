package processors

import (
	"context"
	"fmt"
	"xvitu/sec-bot/domain"
	domainEntity "xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/entypoint/dto"
	"xvitu/sec-bot/infra/persistence/repository"
	"xvitu/sec-bot/infra/telegram"
)

type FaqProcessor struct {
	chatRepository    repository.ChatRepositoryInterface
	telegramGateway   *telegram.Gateway
	messageRepository repository.MessageRepositoryInterface
}

const backMessageId = "11"

func NewFaqProcessor(
	chatRepository repository.ChatRepositoryInterface,
	telegramGateway *telegram.Gateway,
	messageRepository repository.MessageRepositoryInterface,
) *FaqProcessor {
	return &FaqProcessor{
		chatRepository:    chatRepository,
		telegramGateway:   telegramGateway,
		messageRepository: messageRepository,
	}
}

func (p *FaqProcessor) Execute(chatUpdate dto.Chat, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	if chatUpdate.Message == backMessageId {
		return p.handleMessage(domain.Start, "greetings", chat)
	}

	chat, chatError := p.handleMessage(domain.Faq, "faq_"+chatUpdate.Message, chat)
	if chatError != nil {
		return p.handleMessage(domain.Faq, "invalid_option", chat)
	}

	return chat, nil
}

func (p *FaqProcessor) handleMessage(step domain.Step, messageId string, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	replyMessage := p.messageRepository.GetByStepAndMessageId(step, messageId)
	if replyMessage == nil {
		return nil, fmt.Errorf("no Message Found")
	}

	_, sendMessageError := p.telegramGateway.SendMessage(chat.ExternalId, replyMessage.Text)
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
