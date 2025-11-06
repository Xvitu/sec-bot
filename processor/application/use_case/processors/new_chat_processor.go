package processors

import (
	"github.com/xvitu/sec-bot/processor/application/service"
	"github.com/xvitu/sec-bot/processor/domain"
	domainEntity "github.com/xvitu/sec-bot/processor/domain/entity"
	"github.com/xvitu/sec-bot/processor/entypoint/dto"
	"time"

	"github.com/google/uuid"
)

type NewChatProcessor struct {
	chatService *service.ChatService
}

func CreateNewChatProcessor(chatService *service.ChatService) *NewChatProcessor {
	return &NewChatProcessor{chatService: chatService}
}

const InitialMessageId = "greetings"

func (p *NewChatProcessor) Execute(chatUpdate dto.Chat, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	domainChat := &domainEntity.Chat{
		Origin:        chatUpdate.Origin,
		ExternalId:    chatUpdate.ExternalId,
		UserId:        chatUpdate.ExternalUserId,
		CreatedAt:     time.Now().String(),
		UpdatedAt:     time.Now().String(),
		Step:          domain.MainMenu,
		LastMessageID: InitialMessageId,
		Id:            uuid.New().String(),
	}

	return p.chatService.HandleReplyMessages(domain.MainMenu, []string{InitialMessageId}, domainChat)
}
