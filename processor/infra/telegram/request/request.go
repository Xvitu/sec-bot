package request

type TeleGramRequest interface {
	Endpoint() string
	Body() ([]byte, error)
}
