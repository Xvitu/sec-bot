package processors

import (
	"math/rand"

	"github.com/xvitu/sec-bot/processor/application/service"
	"github.com/xvitu/sec-bot/processor/domain"
	domainEntity "github.com/xvitu/sec-bot/processor/domain/entity"
	"github.com/xvitu/sec-bot/processor/entypoint/dto"
	"github.com/xvitu/sec-bot/processor/infra/persistence/repository"
)

type TipsProcessor struct {
	chatService       *service.ChatService
	messageRepository repository.MessageRepositoryInterface
}

func NewTipsProcessor(
	chatService *service.ChatService,
	messageRepository repository.MessageRepositoryInterface,
) *TipsProcessor {
	return &TipsProcessor{
		chatService:       chatService,
		messageRepository: messageRepository,
	}
}

const (
	Back     = "2"
	MoreTips = "1"
	TipMenu  = "tip_menu"
)

func (p *TipsProcessor) Execute(chatUpdate *dto.Chat, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	var messages []string
	var step domain.Step

	switch chatUpdate.Message {
	case MoreTips:
		tip := p.randomTip(chat)
		messages = []string{tip.Id, TipMenu}
		step = domain.Tips
		break
	case Back:
		messages = []string{"greetings"}
		step = domain.MainMenu
		break
	default:
		return p.chatService.HandleError("invalid_option", chat)
	}

	return p.chatService.HandleReplyMessages(step, messages, chat)
}

func (p *TipsProcessor) randomTip(chat *domainEntity.Chat) *domainEntity.Message {
	lastMessageId := chat.LastMessageID
	allMessages := p.messageRepository.FindAllByStepExcludingIds(domain.Tips, []string{lastMessageId, TipMenu})
	return allMessages[rand.Intn(len(allMessages))]
}
