package request

type ChatUpdateRequest struct {
	UpdateID int64 `json:"update_id"`
	Message  struct {
		MessageID int64  `json:"message_id"`
		Text      string `json:"text"`
		Chat      struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}
