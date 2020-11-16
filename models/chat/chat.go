package chat 

import (
	"sort"
	"sync"
)

/*Chat hold messages and info about chat*/
type Chat struct {
	ChatID int
	Messages []Message
}

/*NewChat creates a new chat. You can pass in '-1' as a default value for chatID or sessionID which will not set the structs ID*/
func NewChat(chatID int, messages []Message) *Chat {
	chat := new(Chat)
	
	if chatID < 0 {
		chat.ChatID = chatID
	}

	if len(messages) > 0 {
		chat.Messages = messages
	} else {
		chat.Messages = make([]Message, 10) //default length of messages to 10
	}

	return chat
}

/*AddMessage append messaage to Messages slice. 
	Sorts messages by timestamp. Recommended to run as a routine*/
func (chat *Chat) AddMessage(message Message, wg *sync.WaitGroup) {
	chat.Messages = append(chat.Messages, message)
	wg.Done()
}

/*SortMessages sort messages in a chat by timestamp*/
func (chat *Chat) SortMessages(wg *sync.WaitGroup) {
	/*Keep messages sorted by time stamp*/
	sort.Slice(chat.Messages, func(p, q int) bool {
		return chat.Messages[p].Timestamp < chat.Messages[q].Timestamp
	})
	wg.Done()
}
