package chat

import (
	"os"
	"testing"
	"time"
)

var chatRepo *Repo
var chatID int = 1
var timestamp int64 = time.Now().Unix()

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	var err error
	chatRepo, err = NewRepo()
	if err != nil {
		panic(err.Error())
	}
}

func teardown() {
	err := chatRepo.Close()
	if err != nil {
		panic(err.Error())
	}
}

func TestNewRepo(t *testing.T) {
	err := chatRepo.DB.Ping()
	if err != nil {
		t.Errorf("Failed to create chatRepo with Database connection! %s", err.Error())
	}
}

func TestAddMessage(t *testing.T) {
	msgID, err := chatRepo.AddMessage(chatID, timestamp, 2, "This is a test message")
	if msgID < 0 || err != nil {
		t.Errorf("Error inserting single new message into DB, %s", err.Error())
	}
}

func TestGetAllMessages(t *testing.T) {
	newChat, err := chatRepo.GetAllMessages(chatID)
	if err != nil {
		t.Errorf("Error getting all messages from DB, %s", err.Error())
	}

	if len(newChat.Messages) == 0 {
		t.Errorf("No messages found for chat, %s", err.Error())
	}
}

func TestGetMessagesSince(t *testing.T) {
	newChat, err := chatRepo.GetMessagesSince(timestamp, chatID)
	if err != nil {
		t.Errorf("Error getting messages since %v from DB, %s", timestamp, err.Error())
	}

	messages := newChat.Messages
	if len(messages) == 0 {
		t.Errorf("No messages found for chat, %s", err.Error())
	}
}

func TestRemoveChatMessages(t *testing.T) {
	err := chatRepo.RemoveChatMessages(chatID)
	if err != nil {
		t.Errorf("Error removing chat messages from DB, %s", err.Error())
	}
}
