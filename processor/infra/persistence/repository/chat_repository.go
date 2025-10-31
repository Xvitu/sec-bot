package repository

import (
	"context"
	"fmt"
	domain "xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/infra/persistence/entity"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ChatRepositoryInterface interface {
	Save(ctx context.Context, chat domain.Chat) error
	FindByExternalId(ctx context.Context, id string) (*domain.Chat, error)
}

type ChatRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewChatRepository(client *dynamodb.Client) *ChatRepository {
	return &ChatRepository{
		client:    client,
		tableName: "Chats",
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

// todo - ver questao de custo do GSI
func (r *ChatRepository) FindByExternalId(ctx context.Context, id string) (*domain.Chat, error) {
	out, err := r.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              &r.tableName,
		IndexName:              aws.String("ExternalIDIndex"),
		KeyConditionExpression: aws.String("external_id = :eid"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":eid": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar chat: %w", err)
	}

	if len(out.Items) == 0 {
		return nil, nil
	}

	item := out.Items[0]
	var chatEntity entity.Chat
	if err := attributevalue.UnmarshalMap(item, &chatEntity); err != nil {
		return nil, fmt.Errorf("erro ao deserializar chat: %w", err)
	}

	domainChat := entity.ToDomain(chatEntity)
	return &domainChat, nil
}
