package processors

import (
	"math/rand"
	"xvitu/sec-bot/application/service"
	"xvitu/sec-bot/domain"
	domainEntity "xvitu/sec-bot/domain/entity"
	"xvitu/sec-bot/entypoint/dto"
	"xvitu/sec-bot/infra/persistence/repository"
)

type QuizProcessor struct {
	chatService       *service.ChatService
	messageRepository repository.MessageRepositoryInterface
}

func NewQuizProcessor(
	chatService *service.ChatService,
	messageRepository repository.MessageRepositoryInterface,
) *QuizProcessor {
	return &QuizProcessor{chatService: chatService, messageRepository: messageRepository}
}

const QuizBack = "2"
const SendQuestion = "1"
const QuizMenu = "quiz_menu"

func (p *QuizProcessor) Execute(chatUpdate dto.Chat, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
	var messages []string
	var step domain.Step
	lastMessage := chat.LastMessageID

	if lastMessage == QuizMenu {
		switch chatUpdate.Message {
		case SendQuestion:
			quiz := p.randomQuestion(chat)
			messages = []string{quiz.Id}
			step = domain.QuizQuestion
			break
		case QuizBack:
			messages = []string{"greetings"}
			step = domain.MainMenu
			break
		default:
			return p.chatService.HandleError("invalid_option", chat)
		}

		return p.chatService.HandleReplyMessages(step, messages, chat)
	}

	answer := p.findAnswer(lastMessage)
	if answer.Text == chatUpdate.Message {
		p.chatService.HandleReplyMessages(domain.QuizFeedback, []string{"quiz_success"}, chat)
		return p.chatService.HandleReplyMessages(domain.Quiz, []string{QuizMenu}, chat)
	}

	p.chatService.HandleReplyMessages(domain.QuizFeedback, []string{"quiz_error"}, chat)
	p.chatService.HandleReplyMessages(domain.QuizExplanation, []string{lastMessage}, chat)

	return p.chatService.HandleReplyMessages(domain.Quiz, []string{QuizMenu}, chat)
}

func (p *QuizProcessor) randomQuestion(chat *domainEntity.Chat) *domainEntity.Message {
	lastMessageId := chat.LastMessageID
	allMessages := p.messageRepository.FindAllByStepExcludingIds(domain.QuizQuestion, []string{lastMessageId})
	return allMessages[rand.Intn(len(allMessages))]
}

func (p *QuizProcessor) findAnswer(lastMessageId string) *domainEntity.Message {
	return p.messageRepository.GetByStepAndMessageId(domain.QuizAnswer, lastMessageId)
}
