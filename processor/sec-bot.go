package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"xvitu/sec-bot/application/service"
	"xvitu/sec-bot/application/use_case"
	"xvitu/sec-bot/application/use_case/processors"
	"xvitu/sec-bot/domain"
	"xvitu/sec-bot/entypoint/dto"
	telegramClient "xvitu/sec-bot/infra/client/telegram"
	telegramGateway "xvitu/sec-bot/infra/gateway/communication/telegram"
	"xvitu/sec-bot/infra/persistence/mongodb"
	"xvitu/sec-bot/infra/persistence/repository"
	"xvitu/sec-bot/shared/env"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		envs := env.Get()

		mongoCli := mongodb.NewClient(envs.MongoUrl, envs.DbName).Connect()

		telegramCli := telegramClient.NewTelegramClient(envs)
		chatRepository := repository.NewChatRepository(mongoCli)
		gateway := telegramGateway.NewGateway(telegramCli)
		messageRepository := &repository.MessageRepository{}

		chatService := service.NewChatService(chatRepository, gateway, messageRepository)

		quizProcessor := processors.NewQuizProcessor(chatService, messageRepository)

		useCase := use_case.NewChatUpdateHandler(
			map[domain.Step]processors.MessageProcessor{
				domain.Start:           processors.CreateNewChatProcessor(chatService),
				domain.Faq:             processors.NewFaqProcessor(chatService),
				domain.MainMenu:        processors.NewMainMenuProcessor(chatService),
				domain.Tips:            processors.NewTipsProcessor(chatService, messageRepository),
				domain.Scams:           processors.NewScamProcessor(chatService),
				domain.QuizExplanation: quizProcessor,
				domain.Quiz:            quizProcessor,
				domain.QuizAnswer:      quizProcessor,
				domain.QuizQuestion:    quizProcessor,
				domain.QuizFeedback:    quizProcessor,
			},
			chatRepository,
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
