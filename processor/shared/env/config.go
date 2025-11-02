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
	MongoUrl              string
	AwsRegion             string
	DbName                string
}

var (
	instance *Config
	once     sync.Once
)

func Get() *Config {
	once.Do(func() {
		_ = godotenv.Load()

		instance = &Config{
			TelegramBotToken:      os.Getenv("TELEGRAM_BOT_TOKEN"),
			TelegramUrl:           os.Getenv("TELEGRAM_URL"),
			Env:                   os.Getenv("ENV"),
			TeleGramClientTimeout: os.Getenv("TELEGRAM_CLIENT_TIMEOUT"),
			MongoUrl:              os.Getenv("MONGO_URL"),
			AwsRegion:             os.Getenv("AWS_REGION"),
			DbName:                os.Getenv("DB_NAME"),
		}
	})
	return instance
}
