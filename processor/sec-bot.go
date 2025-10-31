package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"xvitu/sec-bot/domain"
	"xvitu/sec-bot/entypoint/dto"
	"xvitu/sec-bot/infra/persistence/dynamo"
	"xvitu/sec-bot/infra/persistence/repository"
	"xvitu/sec-bot/infra/telegram"
	"xvitu/sec-bot/shared/env"
	"xvitu/sec-bot/use_case"
)

func main() {

	// todo - rever migrations
	//ctx := context.TODO()
	//dynamoClient, _ := dynamo.NewClient(context.TODO())
	//err := dynamo.EnsureTableExists(ctx, dynamoClient, "Chats")
	//if err != nil {
	//	log.Fatal(err)
	//}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dynamoCli, _ := dynamo.NewClient(context.TODO(), env.Get())
		telegramCli := telegram.NewTelegramClient(env.Get())
		useCase := use_case.NewChatUpdateProcessor(
			repository.NewChatRepository(dynamoCli),
			telegram.NewGateway(telegramCli),
			&repository.MessageRepository{},
		)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "erro ao ler body", http.StatusBadRequest)
			return
		}

		var data map[string]interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			http.Error(w, "erro ao fazer parse do JSON", http.StatusBadRequest)
			return
		}
		message, _ := data["message"].(string)

		chatDto := dto.Chat{
			ExternalId:     "5470945009",
			ExternalUserId: "5470945009",
			SentAt:         time.Now().String(),
			Message:        message,
			Origin:         domain.Telegram,
		}

		_, e := useCase.Run(chatDto)

		if e != nil {
		}

	})

	fmt.Println("Running bot on port 8080")
	http.ListenAndServe(":8080", nil)
}
