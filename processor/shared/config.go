package shared

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBotToken string
	TelegramUrl      string
	Env              string
}

var (
	instance *Config
	once     sync.Once
)

func Get() *Config {
	once.Do(func() {
		_ = godotenv.Load()

		instance = &Config{
			TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
			TelegramUrl:      os.Getenv("TELEGRAM_URL"),
			Env:              os.Getenv("ENV"),
		}
	})
	return instance
}
