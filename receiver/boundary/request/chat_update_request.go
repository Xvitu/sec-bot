package request

type ChatUpdateRequest struct {
	UpdateID int64 `json:"update_id"`
	Message  struct {
		MessageID int64  `json:"message_id"`
		Text      string `json:"text"`
		Chat      struct {
			Id int64 `json:"id"`
		} `json:"chat"`
		From struct {
			Id        int64  `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
		} `json:"from"`
	} `json:"message"`
}
