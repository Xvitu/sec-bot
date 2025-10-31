package processors

import (
	"context"
	"xvitu/sec-bot/domain"
	domainEntity "xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/entypoint/dto"
	"xvitu/sec-bot/infra/persistence/repository"
	"xvitu/sec-bot/infra/telegram"
)

type MainMenuProcessor struct {
	chatRepository    repository.ChatRepositoryInterface
	telegramGateway   *telegram.Gateway
	messageRepository repository.MessageRepositoryInterface
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

func NewMainMenuProcessor(
	chatRepository repository.ChatRepositoryInterface,
	telegramGateway *telegram.Gateway,
	messageRepository repository.MessageRepositoryInterface,
) *MainMenuProcessor {
	return &MainMenuProcessor{
		messageRepository: messageRepository,
		chatRepository:    chatRepository,
		telegramGateway:   telegramGateway,
	}
}

// todo - rever classe e parse do step e menu opt
// todo - implementar strategy
func (p *MainMenuProcessor) Execute(chatUpdate dto.Chat, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	option := MenuOptions{chatUpdate.Message}
	var replyMessage *domainEntity.Message
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
		messageId = "faq_menu" // todo - random message
		break
	case Scams:
		step = domain.Scams
		messageId = "faq_menu" // todo - mudar para menu
		break
	default:
		step = domain.Start
		messageId = "invalid_option"
		break
	}

	replyMessage = p.messageRepository.GetByStepAndMessageId(step, messageId)

	p.telegramGateway.SendMessage(chat.ExternalId, replyMessage.Text)

	chat.Step = step
	chat.LastMessageID = messageId
	p.chatRepository.Save(context.Background(), *chat)
	return chat, nil
}
