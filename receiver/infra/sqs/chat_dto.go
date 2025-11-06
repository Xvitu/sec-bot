package sqs

import "github.com/xvitu/sec-bot/receiver/domain"

type Chat struct {
	ExternalId     string        `json:"external_id"`
	ExternalUserId string        `json:"external_user_id"`
	SentAt         string        `json:"created_at"`
	Message        string        `json:"message"`
	Origin         domain.Origin `json:"origin"`
}
