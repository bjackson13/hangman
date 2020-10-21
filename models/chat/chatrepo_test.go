package chat

import (
	"testing"
	"os"
	"time"
)

var chatRepo *Repo
var chatID int
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

func TestAddChat(t *testing.T) {
	insertedChatID, err := chatRepo.AddChat()
	if err != nil {
		t.Errorf("Error inserting new chat into DB, %s", err.Error())
	}

	chatID = insertedChatID	
}

func TestAddChatUsers(t *testing.T) {
	err := chatRepo.AddChatUsers(chatID, 2, 3) //test users
	if err != nil {
		t.Errorf("Error inserting new chat users into DB, %s", err.Error())
	}
}

func TestAddMessage(t *testing.T) {
	msgID, err := chatRepo.AddMessage(chatID, timestamp, 1, "This is a test message")
	if msgID < 0 || err != nil {
		t.Errorf("Error inserting single new message into DB, %s", err.Error())
	}
}

func TestGetAllMessages(t *testing.T) {
	newChat, err := chatRepo.GetAllMessages(chatID)
	if err != nil {
		t.Errorf("Error getting all messages from DB, %s", err.Error())
	}

	messages := newChat.Messages
	if len(messages) == 0 {
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

func TestRemoveChatUsers(t *testing.T) {
	err := chatRepo.RemoveChatUsers(chatID)
	if err != nil {
		t.Errorf("Error removing chat users from DB, %s", err.Error())
	}
}

func TestRemoveChatMessages(t *testing.T) {
	err := chatRepo.RemoveChatMessages(chatID)
	if err != nil {
		t.Errorf("Error removing chat messages from DB, %s", err.Error())
	}
}

func TestRemoveChat(t *testing.T) {
	err := chatRepo.RemoveChat(chatID)
	if err != nil {
		t.Errorf("Error removing chat from DB, %s", err.Error())
	}
}
