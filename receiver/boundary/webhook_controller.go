package boundary

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xvitu/sec-bot/receiver/boundary/request"
)

type WebHookController struct{}

func (c *WebHookController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	var update request.ChatUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "JSON inválido", http.StatusUnprocessableEntity)
		return
	}

	// todo - postar na fila
	fmt.Printf("Mensagem recebida: %s (Chat ID: %d)\n", update.Message.Text, update.Message.Chat.ID)

	w.WriteHeader(http.StatusOK)
}
