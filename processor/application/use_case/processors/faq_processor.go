package processors

import (
	"github.com/xvitu/sec-bot/processor/application/service"
	"github.com/xvitu/sec-bot/processor/domain"
	domainEntity "github.com/xvitu/sec-bot/processor/domain/entity"
	"github.com/xvitu/sec-bot/processor/entypoint/dto"
)

type FaqProcessor struct {
	chatService *service.ChatService
}

const faqBackMessageId = "11"

func NewFaqProcessor(
	chatService *service.ChatService,
) *FaqProcessor {
	return &FaqProcessor{chatService: chatService}
}

func (p *FaqProcessor) Execute(chatUpdate dto.Chat, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	if chatUpdate.Message == faqBackMessageId {
		return p.chatService.HandleReplyMessages(domain.MainMenu, []string{"greetings"}, chat)
	}

	messageId := "faq_" + chatUpdate.Message
	updatedChat, chatError := p.chatService.HandleReplyMessages(domain.Faq, []string{messageId}, chat)
	if chatError != nil {
		return p.chatService.HandleError("invalid_option", chat)
	}

	return updatedChat, nil
}
