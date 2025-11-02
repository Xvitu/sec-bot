package processors

import (
	"xvitu/sec-bot/application/service"
	"xvitu/sec-bot/domain"
	domainEntity "xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/entypoint/dto"
)

type FaqProcessor struct {
	chatService *service.ChatService
}

const backMessageId = "11"

func NewFaqProcessor(
	chatService *service.ChatService,
) *FaqProcessor {
	return &FaqProcessor{chatService: chatService}
}

func (p *FaqProcessor) Execute(chatUpdate dto.Chat, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	if chatUpdate.Message == backMessageId {
		return p.chatService.HandleReplyMessages(domain.MainMenu, []string{"greetings"}, chat)
	}

	messageId := "faq_" + chatUpdate.Message
	updatedChat, chatError := p.chatService.HandleReplyMessages(domain.Faq, []string{messageId}, chat)
	if chatError != nil {
		return p.chatService.HandleError("invalid_option", chat)
	}

	return updatedChat, nil
}
