package response

type Chat struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}

type Message struct {
	MessageID int64  `json:"message_id"`
	Chat      Chat   `json:"chat"`
	Text      string `json:"text"`
}

type SendMessageResponse struct {
	RawMessage []byte  `json:"raw_message"`
	Result     Message `json:"result"`
}
