package use_case

import (
	"context"
	"time"
	"xvitu/sec-bot/domain"
	domainEntity "xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/entypoint/dto"
	"xvitu/sec-bot/infra/persistence/repository"
	"xvitu/sec-bot/infra/telegram"

	"github.com/google/uuid"
)

type ChatUpdateProcessor struct {
	chatRepository    *repository.ChatRepository
	telegramGateway   *telegram.Gateway
	messageRepository *repository.MessageRepository
}

func NewChatUpdateProcessor(
	chatRepository *repository.ChatRepository,
	telegramGateway *telegram.Gateway,
	messageRepository *repository.MessageRepository,
) *ChatUpdateProcessor {
	return &ChatUpdateProcessor{
		chatRepository:    chatRepository,
		telegramGateway:   telegramGateway,
		messageRepository: messageRepository,
	}
}

func (u *ChatUpdateProcessor) Run(chatUpdate dto.Chat) (*domainEntity.Chat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	persistedChat, chatRepoError := u.chatRepository.FindByExternalId(ctx, chatUpdate.ExternalUserId)

	if chatRepoError != nil {
		return nil, chatRepoError
	}

	if persistedChat == nil {
		message := u.messageRepository.GetByStepAndMessageId(domain.Start, "greetings")
		_, sendMessageError := u.telegramGateway.SendMessage(chatUpdate.ExternalId, message.Text)

		if sendMessageError != nil {
			return nil, sendMessageError
		}

		domainChat := domainEntity.Chat{
			Origin:        chatUpdate.Origin,
			ExternalId:    chatUpdate.ExternalId,
			UserId:        chatUpdate.ExternalUserId,
			CreatedAt:     time.Now().String(),
			UpdatedAt:     time.Now().String(),
			Step:          domain.Start,
			LastMessageID: message.Id,
			Id:            uuid.New().String(),
		}
		createErr := u.chatRepository.Save(context.Background(), domainChat)

		if createErr != nil {
			return nil, createErr
		}

		return &domainChat, nil
	} else {
		// todo - rever tudo qui
		baseProcessor := BaseProcessor{
			chatRepository:    u.chatRepository,
			telegramGateway:   u.telegramGateway,
			messageRepository: u.messageRepository,
		}
		baseProcessor.execute(chatUpdate, persistedChat)
	}

	return nil, nil
}

type BaseProcessor struct {
	chatRepository    *repository.ChatRepository
	telegramGateway   *telegram.Gateway
	messageRepository *repository.MessageRepository
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

// todo - rever classe e parse do step e menu opt
func (p *BaseProcessor) execute(chatUpdate dto.Chat, chat *domainEntity.Chat) (*domainEntity.Chat, error) {

	option := MenuOptions{chatUpdate.Message}
	var replyMessage domainEntity.Message
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
		messageId = "faq_menu" // todo - random message
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
