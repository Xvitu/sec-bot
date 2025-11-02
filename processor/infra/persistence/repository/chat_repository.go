package repository

import (
	"context"
	"errors"
	"fmt"
	domain "xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/infra/persistence/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatRepositoryInterface interface {
	Save(ctx context.Context, chat domain.Chat) error
	FindByExternalId(ctx context.Context, id string) (*domain.Chat, error)
}

type ChatRepository struct {
	database  *mongo.Database
	tableName string
}

func NewChatRepository(database *mongo.Database) *ChatRepository {
	return &ChatRepository{
		database:  database,
		tableName: "Chats",
	}
}

func (r *ChatRepository) Save(ctx context.Context, chat domain.Chat) error {
	chatEntity := entity.FromDomain(chat)

	collection := r.database.Collection(r.tableName)

	_, err := collection.InsertOne(ctx, chatEntity)
	if err != nil {
		return fmt.Errorf("erro ao salvar chat no MongoDB: %w", err)
	}

	return nil
}

func (r *ChatRepository) FindByExternalId(ctx context.Context, id string) (*domain.Chat, error) {
	collection := r.database.Collection(r.tableName)

	var chatEntity entity.Chat
	err := collection.FindOne(ctx, bson.M{"external_id": id}).Decode(&chatEntity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, fmt.Errorf("erro ao buscar chat: %w", err)
	}

	domainChat := entity.ToDomain(chatEntity)
	return &domainChat, nil
}
