package processors

import (
	domainEntity "github.com/xvitu/sec-bot/processor/domain/entity"
	"github.com/xvitu/sec-bot/processor/entypoint/dto"
)

type MessageProcessor interface {
	Execute(chatUpdate *dto.Chat, chat *domainEntity.Chat) (*domainEntity.Chat, error)
}
