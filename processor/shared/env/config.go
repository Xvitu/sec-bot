package env

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBotToken      string
	TelegramUrl           string
	Env                   string
	TeleGramClientTimeout string
	DynamoDBEndpoint      string
	AwsRegion             string
}

var (
	instance *Config
	once     sync.Once
)

func Get() *Config {
	once.Do(func() {
		_ = godotenv.Load("./processor/.env")

		instance = &Config{
			TelegramBotToken:      os.Getenv("TELEGRAM_BOT_TOKEN"),
			TelegramUrl:           os.Getenv("TELEGRAM_URL"),
			Env:                   os.Getenv("ENV"),
			TeleGramClientTimeout: os.Getenv("TELEGRAM_CLIENT_TIMEOUT"),
			DynamoDBEndpoint:      os.Getenv("DYNAMODB_ENDPOINT"),
			AwsRegion:             os.Getenv("AWS_REGION"),
		}
	})
	return instance
}
