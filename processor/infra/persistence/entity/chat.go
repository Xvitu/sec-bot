package entity

import domain "xvitu/sec-bot/domain/entity"

type Chat struct {
	Id            int    `dynamodbav:"id"`
	ExternalId    string `dynamodbav:"external_id"`
	UserId        string `dynamodbav:"user_id"`
	LastMessageID string `dynamodbav:"last_message_id"`
	CreatedAt     string `dynamodbav:"created_at"`
	UpdatedAt     string `dynamodbav:"updated_at"`
}

func FromDomain(chat domain.Chat) Chat {
	return Chat{
		Id:            chat.Id,
		ExternalId:    chat.ExternalId,
		UserId:        chat.UserId,
		LastMessageID: chat.LastMessageID,
		CreatedAt:     chat.CreatedAt,
		UpdatedAt:     chat.UpdatedAt,
	}
}

func ToDomain(chat Chat) domain.Chat {
	return domain.Chat{
		Id:            chat.Id,
		ExternalId:    chat.ExternalId,
		UserId:        chat.UserId,
		LastMessageID: chat.LastMessageID,
		CreatedAt:     chat.CreatedAt,
		UpdatedAt:     chat.UpdatedAt,
	}
}
