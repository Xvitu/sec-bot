package use_case

type ChatService struct {
}

//func (s *ChatService) handleError(messageId string, chat *domainEntity.Chat) (*domainEntity.Chat, error) {
//	replyMessage := s.messageRepository.GetByStepAndMessageId(domain.Error, messageId)
//	//_, sendMessageError := s.telegramGateway.SendMessage(chat.ExternalId, replyMessage.Text)
//	if sendMessageError != nil {
//		return nil, fmt.Errorf("error while sending message: %s", sendMessageError)
//	}
//
//	return chat, nil
//}
