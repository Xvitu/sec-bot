package response

type Chat struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Message struct {
	MessageID string `json:"message_id"`
	Chat      Chat   `json:"chat"`
	Text      string `json:"text"`
}

type SendMessageResponse struct {
	Result Message `json:"result"`
}
