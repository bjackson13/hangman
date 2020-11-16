package chat

import (
	"github.com/bjackson13/hangman/models/chat"
	"strconv"
)

/*Service struct to bind our service functions to*/
type Service struct {}

/*NewService produce a new service*/
func NewService() *Service {
	return new(Service)
}

/*GetAllMessages returns all messages for a given chat*/
func (s *Service) GetAllMessages(chatID int) (*chat.Chat, error) {
	chatRepo, err := chat.NewRepo()
	defer chatRepo.Close()
	if err != nil {
		return nil, err
	}

	return chatRepo.GetAllMessages(chatID)
}

/*GetMessagesSince returns all messages since a certain point in time*/
func (s *Service) GetMessagesSince(time string, chatID int) (*chat.Chat, error) {
	chatRepo, err := chat.NewRepo()
	defer chatRepo.Close()
	if err != nil {
		return nil, err
	}

	parsedTime, err := strconv.ParseInt(time, 10, 64)
	if err != nil {
		return nil, err
	}

	return chatRepo.GetMessagesSince(parsedTime, chatID)
}

/*AddMessage add a messasge to a chat*/
func (s *Service) AddMessage(chatID, senderID int, time int64, text string) error {
	chatRepo, err := chat.NewRepo()
	defer chatRepo.Close()
	if err != nil {
		return err
	}
	_, err = chatRepo.AddMessage(chatID, time, senderID, text)
	return err
}