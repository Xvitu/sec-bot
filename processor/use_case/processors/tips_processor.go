package processors

import (
	domainEntity "xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/entypoint/dto"
	"xvitu/sec-bot/infra/persistence/repository"
	"xvitu/sec-bot/infra/telegram"
)

type TipsProcessor struct {
	chatRepository    repository.ChatRepositoryInterface
	telegramGateway   *telegram.Gateway
	messageRepository repository.MessageRepositoryInterface
}

func NewTipsProcessor(
	chatRepository repository.ChatRepositoryInterface,
	telegramGateway *telegram.Gateway,
	messageRepository repository.MessageRepositoryInterface,
) *TipsProcessor {
	return &TipsProcessor{
		chatRepository:    chatRepository,
		telegramGateway:   telegramGateway,
		messageRepository: messageRepository,
	}
}

const (
	Back     = "1"
	MoreTips = "2"
)

func (p *TipsProcessor) Execute(chatUpdate dto.Chat, chat *domainEntity.Chat) (*domainEntity.Chat, error) {

}

func (p *TipsProcessor) randomTip() {

}
