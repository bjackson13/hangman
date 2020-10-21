package chat

import (
	"github.com/bjackson13/hangman/models"
	"database/sql"
	"sync"
)

/*Repo - Struct for CRUDing users from the database*/
type Repo struct {
	DB *sql.DB
}

/*NewRepo - Create new repo with acce3ss to mysql database*/
func NewRepo() (*Repo, error) {
	conn, err := dbconn.Connect() 
	if err != nil {
		return nil, err
	}

	repo := new(Repo)
	repo.DB = conn
	return repo, nil
}

/*Close closes the database connection*/
func (repo *Repo) Close() error {
	return repo.DB.Close()
	
}

/*GetAllMessages - get all messages from the database for a given chat*/
func (repo *Repo) GetAllMessages(chatID int) (*Chat, error) {
	conn := repo.DB
	chat := NewChat(chatID, nil)
	msgStmt, err := conn.Prepare("SELECT Chat.ChatId, MessageId, Timestamp, SenderId, MessageText FROM Messages JOIN Chat ON Messages.ChatId WHERE Chat.ChatId = ?")
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
	conn := repo.DB
	chat := NewChat(chatID, nil)
	msgStmt, err := conn.Prepare("SELECT Chat.ChatId, MessageId, Timestamp, SenderId, MessageText FROM Messages JOIN Chat ON Messages.ChatId WHERE Chat.ChatId = ? AND Timestamp >= ?")
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
	conn := repo.DB
	chatStmt, err := conn.Prepare("INSERT INTO Messages(ChatID, Timestamp, SenderId, MessageText) VALUES (?,?,?,?)")
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

/*AddChat make a new chat*/
func (repo *Repo) AddChat() (int, error) {
	conn := repo.DB
	chatStmt, err := conn.Prepare("INSERT INTO Chat VALUE()")
	if err != nil {
		return -1, err
	}
	defer chatStmt.Close()

	res, err := chatStmt.Exec()
	if err != nil {
		return -1, err
	}
	lastID, err := res.LastInsertId()
	return int(lastID), err
}

/*AddSingleChatUser add a single user to a chat*/
func (repo *Repo) AddSingleChatUser(chatID int, user int) error {
	conn := repo.DB
	chatStmt, err := conn.Prepare("INSERT INTO ChatUsers(UserId, ChatId) VALUES (?,?)")
	if err != nil {
		return err
	}
	defer chatStmt.Close()

	_, err = chatStmt.Exec(user, chatID)
	if err != nil {
		return err
	}
	return err
}

/*AddChatUsers add users to a chat*/
func (repo *Repo) AddChatUsers(chatID int, user1 int, user2 int) error {
	conn := repo.DB
	chatStmt, err := conn.Prepare("INSERT INTO ChatUsers(UserId, ChatId) VALUES (?,?), (?,?)")
	if err != nil {
		return err
	}
	defer chatStmt.Close()

	_, err = chatStmt.Exec(user1, chatID, user2, chatID)
	if err != nil {
		return err
	}
	return err
}

/*RemoveChatUsers remove all users from a chat*/
func (repo *Repo) RemoveChatUsers(chatID int) error {
	conn := repo.DB
	chatStmt, err := conn.Prepare("DELETE FROM ChatUsers WHERE ChatId = ?")
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

/*RemoveChatMessages remove all messages from a chat*/
func (repo *Repo) RemoveChatMessages(chatID int) error {
	conn := repo.DB
	chatStmt, err := conn.Prepare("DELETE FROM Messages WHERE ChatId = ?")
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

/*RemoveChat remove a chat from Chat table by ChatID*/
func (repo *Repo) RemoveChat(chatID int) error {
	conn := repo.DB
	chatStmt, err := conn.Prepare("DELETE FROM Chat WHERE ChatId = ?")
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
