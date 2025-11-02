package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	dbUrl  string
	dbName string
}

func NewClient(dbUrl, dbName string) *Client {
	return &Client{
		dbUrl:  dbUrl,
		dbName: dbName,
	}
}

func (c *Client) Connect() *mongo.Database {
	clientOpts := options.Client().ApplyURI(c.dbUrl)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Erro ao fazer ping:", err)
	}

	return client.Database(c.dbName)
}
