package entity

import "xvitu/sec-bot/domain"

type Chat struct {
	Id            string        `json:"id"`
	ExternalId    string        `json:"external_id"`
	UserId        string        `json:"user_id"`
	LastMessageID string        `json:"last_message_id"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
	Step          domain.Step   `json:"step"`
	Origin        domain.Origin `json:"origin"`
}
