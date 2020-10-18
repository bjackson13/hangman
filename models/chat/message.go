package chat

/*Message represents each individual message of a chat*/
type Message struct {
	ChatID int
	MessageID int
	Timestamp int64
	SenderID int
	Text string
}

/*NewMessage create a new message*/
func NewMessage(chatID int, messageID int, timestamp int64, senderID int, text string) *Message {
	message := new(Message)
	message.ChatID = chatID
	message.MessageID = messageID
	message.Timestamp = timestamp
	message.SenderID = senderID
	message.Text = text
	return message
}