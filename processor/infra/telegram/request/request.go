package request

type TelegramRequest interface {
	Endpoint() string
	Body() ([]byte, error)
}
