package processors

import (
	"xvitu/sec-bot/application/service"
	"xvitu/sec-bot/domain"
	domainEntity "xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/entypoint/dto"
)

type ScamProcessor struct {
	chatService *service.ChatService
}

const scamBackMessageId = "11"

func NewScamProcessor(chatService *service.ChatService) *ScamProcessor {
	return &ScamProcessor{chatService: chatService}
}

func (p *ScamProcessor) Execute(chatUpdate dto.Chat, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	if chatUpdate.Message == scamBackMessageId {
		return p.chatService.HandleReplyMessages(domain.MainMenu, []string{"greetings"}, chat)
	}

	messageId := "scam_" + chatUpdate.Message
	updatedChat, chatError := p.chatService.HandleReplyMessages(domain.Scams, []string{messageId}, chat)
	if chatError != nil {
		return p.chatService.HandleError("invalid_option", chat)
	}

	return updatedChat, nil
}
