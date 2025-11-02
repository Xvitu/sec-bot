package processors

import (
	"xvitu/sec-bot/application/service"
	"xvitu/sec-bot/domain"
	domainEntity "xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/entypoint/dto"
)

type MainMenuProcessor struct {
	chatService *service.ChatService
}

type MenuOptions struct {
	option string
}

const (
	Faq   = "1"
	Quiz  = "2"
	Tips  = "3"
	Scams = "4"
)

func NewMainMenuProcessor(chatService *service.ChatService) *MainMenuProcessor {
	return &MainMenuProcessor{chatService: chatService}
}

func (p *MainMenuProcessor) Execute(chatUpdate dto.Chat, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	option := MenuOptions{chatUpdate.Message}
	var step domain.Step
	var messageId string
	switch option.option {
	case Faq:
		step = domain.Faq
		messageId = "faq_menu"
		break
	case Quiz:
		step = domain.Quiz
		messageId = "faq_menu" // todo - random message
		break
	case Tips:
		step = domain.Tips
		messageId = "tip_menu"
		break
	case Scams:
		step = domain.Scams
		messageId = "faq_menu" // todo - mudar para menu
		break
	default:
		return p.chatService.HandleError("invalid_option", chat)
	}

	return p.chatService.HandleReplyMessages(step, []string{messageId}, chat)
}
