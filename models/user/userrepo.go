package user 

import (
	"github.com/bjackson13/hangman/models"
)

/*Repo - Struct for CRUDing users from the database*/
type Repo struct {
	db *dbconn.DB
}

/*New - Create new repo with acce3ss to mysql database*/
func New(conn *dbconn.DB) *Repo {
	repo := new(Repo)
	repo.db = conn
	return repo
}

/*GetUser - get a user from the database with provided credentials*/
func (repo *Repo) GetUser(username string, password string) (User, error) {
	
	conn := repo.db.Connection
	userStmt, err := conn.Prepare("SELECT UserId, Username, IP, UserAgent FROM User WHERE username = ? AND password = ?")
	if err != nil {
		panic(err.Error())
	}
	defer userStmt.Close()
	
	var user User
	err = userStmt.QueryRow(username, password).Scan(&user.UserID, &user.Username, &user.IP, &user.UserAgent)

	return user, err
}