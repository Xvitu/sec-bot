package processors

import (
	domainEntity "xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/entypoint/dto"
)

type MessageProcessor interface {
	Execute(chatUpdate dto.Chat, chat *domainEntity.Chat) (*domainEntity.Chat, error)
}
