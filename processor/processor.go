package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/xvitu/sec-bot/processor/application/service"
	"github.com/xvitu/sec-bot/processor/application/use_case"
	"github.com/xvitu/sec-bot/processor/application/use_case/processors"
	"github.com/xvitu/sec-bot/processor/boudary"
	"github.com/xvitu/sec-bot/processor/domain"
	telegramClient "github.com/xvitu/sec-bot/processor/infra/client/telegram"
	telegramGateway "github.com/xvitu/sec-bot/processor/infra/gateway/communication/telegram"
	"github.com/xvitu/sec-bot/processor/infra/persistence/mongodb"
	"github.com/xvitu/sec-bot/processor/infra/persistence/repository"
	"github.com/xvitu/sec-bot/processor/shared/env"
)

func main() {

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

	lambdaProcessor := boudary.NewLambdaProcessor(useCase)

	lambda.Start(lambdaProcessor.Handler)

}
