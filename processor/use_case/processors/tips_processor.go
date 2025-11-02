package processors

import (
	"context"
	"fmt"
	"math/rand"
	"xvitu/sec-bot/domain"
	domainEntity "xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/entypoint/dto"
	"xvitu/sec-bot/infra/gateway/communication"
	"xvitu/sec-bot/infra/persistence/repository"
)

type TipsProcessor struct {
	chatRepository       repository.ChatRepositoryInterface
	communicationGateway communication.CommunicationGateway
	messageRepository    repository.MessageRepositoryInterface
}

func NewTipsProcessor(
	chatRepository repository.ChatRepositoryInterface,
	communicationGateway communication.CommunicationGateway,
	messageRepository repository.MessageRepositoryInterface,
) *TipsProcessor {
	return &TipsProcessor{
		chatRepository:       chatRepository,
		communicationGateway: communicationGateway,
		messageRepository:    messageRepository,
	}
}

const (
	Back     = "2"
	MoreTips = "1"
	TipMenu  = "tip_menu"
)

func (p *TipsProcessor) Execute(chatUpdate dto.Chat, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	switch chatUpdate.Message {
	case MoreTips:
		tip := p.randomTip(chat)
		p.communicationGateway.SendMessage(chat.ExternalId, tip.Text)

		menuMesage := p.messageRepository.GetByStepAndMessageId(domain.Tips, TipMenu)
		p.communicationGateway.SendMessage(chat.ExternalId, menuMesage.Text)

		p.updateChat(tip.Id, domain.Tips, chat)
		break
	case Back:
		backMessage := p.messageRepository.GetByStepAndMessageId(domain.MainMenu, "greetings")
		p.communicationGateway.SendMessage(chat.ExternalId, backMessage.Text)

		p.updateChat(backMessage.Id, domain.MainMenu, chat)
		break
	default:
		return p.handleError("invalid_option", chat)
	}

	return chat, nil
}

func (p *TipsProcessor) updateChat(messageId string, step domain.Step, chat *domainEntity.Chat) {
	chat.UpdateWithRepledMessage(messageId, step)
	p.chatRepository.Save(context.Background(), *chat)
}

func (p *TipsProcessor) randomTip(chat *domainEntity.Chat) *domainEntity.Message {
	lastMessageId := chat.LastMessageID
	allMessages := p.messageRepository.FindAllByStepExcludingIds(domain.Tips, []string{lastMessageId, TipMenu})
	return allMessages[rand.Intn(len(allMessages))]
}

func (p *TipsProcessor) handleError(messageId string, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	replyMessage := p.messageRepository.GetByStepAndMessageId(domain.Error, messageId)
	_, sendMessageError := p.communicationGateway.SendMessage(chat.ExternalId, replyMessage.Text)
	if sendMessageError != nil {
		return nil, fmt.Errorf("error while sending message: %s", sendMessageError)
	}

	return chat, nil
}
