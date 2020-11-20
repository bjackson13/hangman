package chat

/*Message represents each individual message of a chat*/
type Message struct {
	ChatID int
	MessageID int
	Timestamp int64
	Sender string
	Text string
}

/*NewMessage create a new message*/
func NewMessage(chatID int, messageID int, timestamp int64, sender, text string) *Message {
	message := new(Message)
	message.ChatID = chatID
	message.MessageID = messageID
	message.Timestamp = timestamp
	message.Sender = sender
	message.Text = text
	return message
}