package repository

import (
	"context"
	"errors"
	"fmt"
	domain "github.com/xvitu/sec-bot/processor/domain/entity"
	"github.com/xvitu/sec-bot/processor/infra/persistence/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	filter := bson.M{"_id": chatEntity.Id}
	update := bson.M{"$set": chatEntity}
	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("erro ao salvar chat no MongoDB: %w", err)
	}

	return nil
}

func (r *ChatRepository) FindByExternalId(ctx context.Context, id string) (*domain.Chat, error) {
	collection := r.database.Collection(r.tableName)

	var chatEntity entity.Chat
	err := collection.FindOne(ctx, bson.M{"externalid": id}).Decode(&chatEntity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, fmt.Errorf("erro ao buscar chat: %w", err)
	}

	domainChat := entity.ToDomain(chatEntity)
	return &domainChat, nil
}
