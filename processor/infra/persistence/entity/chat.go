package entity

import "xvitu/sec-bot/domain"
import domainEntity "xvitu/sec-bot/domain/entity"

type Chat struct {
	Id            string        `dynamodbav:"id"`
	ExternalId    string        `dynamodbav:"external_id"`
	UserId        string        `dynamodbav:"user_id"`
	LastMessageID string        `dynamodbav:"last_message_id"`
	CreatedAt     string        `dynamodbav:"created_at"`
	UpdatedAt     string        `dynamodbav:"updated_at"`
	Step          domain.Step   `dynamodbav:"step"`
	Origin        domain.Origin `dynamodbav:"origin"`
}

func FromDomain(chat domainEntity.Chat) Chat {
	return Chat{
		Id:            chat.Id,
		ExternalId:    chat.ExternalId,
		UserId:        chat.UserId,
		LastMessageID: chat.LastMessageID,
		CreatedAt:     chat.CreatedAt,
		UpdatedAt:     chat.UpdatedAt,
		Step:          chat.Step,
	}
}

func ToDomain(chat Chat) domainEntity.Chat {
	return domainEntity.Chat{
		Id:            chat.Id,
		ExternalId:    chat.ExternalId,
		UserId:        chat.UserId,
		LastMessageID: chat.LastMessageID,
		CreatedAt:     chat.CreatedAt,
		UpdatedAt:     chat.UpdatedAt,
		Step:          chat.Step,
	}
}
