package repository

import (
	"context"
	"fmt"
	domain "xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/infra/persistence/entity"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ChatRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewChatRepository(client *dynamodb.Client) *ChatRepository {
	return &ChatRepository{
		client:    client,
		tableName: "chats",
	}
}

func (r *ChatRepository) Save(ctx context.Context, chat domain.Chat) error {
	chatEntity := entity.FromDomain(chat)

	item, err := attributevalue.MarshalMap(chatEntity)
	if err != nil {
		return fmt.Errorf("erro ao serializar chat: %w", err)
	}

	_, err = r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      item,
	})
	return err
}

func (r *ChatRepository) FindByID(ctx context.Context, id string) (*domain.Chat, error) {
	out, err := r.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &r.tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar chat: %w", err)
	}

	if out.Item == nil {
		return nil, nil
	}

	var chatEntity entity.Chat
	if err := attributevalue.UnmarshalMap(out.Item, &chatEntity); err != nil {
		return nil, fmt.Errorf("erro ao deserializar chat: %w", err)
	}

	domainChat := entity.ToDomain(chatEntity)
	return &domainChat, nil
}
