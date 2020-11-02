package chat

import (
	"github.com/bjackson13/hangman/models"
	//"database/sql"
	"sync"
)

type Repo struct {
	dbconn.Repo
}

/*NewRepo - Create new repo with access to mysql database*/
func NewRepo() (*Repo, error) {
	conn, err := dbconn.Connect() 
	if err != nil {
		return nil, err
	}

	repo := new(Repo)
	repo.DB = conn
	return repo, nil
}

/*GetAllMessages - get all messages from the database for a given chat*/
func (repo *Repo) GetAllMessages(chatID int) (*Chat, error) {
	chat := NewChat(chatID, nil)
	msgStmt, err := repo.DB.Prepare("SELECT ChatId, MessageId, Timestamp, SenderId, MessageText FROM Messages WHERE ChatId = ?")
	if err != nil {
		return nil, err
	}
	defer msgStmt.Close()
	
	var wg sync.WaitGroup
	rows, err := msgStmt.Query(chatID)
	for rows.Next() {
		msg := Message{}
		err = rows.Scan(&msg.ChatID, &msg.MessageID, &msg.Timestamp, &msg.SenderID, &msg.Text)
		if err != nil {
			break;
		}
		// run as go routine to not lose time to slice appending
		wg.Add(1)
		go chat.AddMessage(msg, &wg)
	}
	wg.Wait()
	return chat, err
}

/*GetMessagesSince - get all messages from the database for a given chat*/
func (repo *Repo) GetMessagesSince(timestamp int64, chatID int) (*Chat, error) {
	chat := NewChat(chatID, nil)
	msgStmt, err := repo.DB.Prepare("SELECT ChatId, MessageId, Timestamp, SenderId, MessageText FROM Messages WHERE ChatId = ? AND Timestamp >= ?")
	if err != nil {
		return nil, err
	}
	defer msgStmt.Close()
	
	var wg sync.WaitGroup
	rows, err := msgStmt.Query(chatID, timestamp)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		msg := Message{}
		err = rows.Scan(&msg.ChatID, &msg.MessageID, &msg.Timestamp, &msg.SenderID, &msg.Text)
		if err != nil {
			break;
		}
		// run as go routine to not lose time to slice appending
		wg.Add(1)
		go chat.AddMessage(msg, &wg)
	}
	wg.Wait()
	return chat, err
}

/*AddMessage add a message to the given chat*/
func (repo *Repo) AddMessage(chatID int, timestamp int64, senderID int, text string) (int, error) {
	chatStmt, err := repo.DB.Prepare("INSERT INTO Messages(ChatID, Timestamp, SenderId, MessageText) VALUES (?,?,?,?)")
	if err != nil {
		return -1, err
	}
	defer chatStmt.Close()

	res, err := chatStmt.Exec(chatID, timestamp, senderID, text)
	if err != nil {
		return -1, err
	}
	lastID, err := res.LastInsertId()
	return int(lastID), err
}

/*RemoveChatMessages remove all messages from a chat*/
func (repo *Repo) RemoveChatMessages(chatID int) error {
	chatStmt, err := repo.DB.Prepare("DELETE FROM Messages WHERE ChatId = ?")
	if err != nil {
		return err
	}
	defer chatStmt.Close()

	_, err = chatStmt.Exec(chatID)
	if err != nil {
		return err
	}
	return err
}
