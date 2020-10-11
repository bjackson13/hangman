package user 

import (
	"github.com/bjackson13/hangman/models"
)

/*
	Struct for CRUDing users from the database
*/
type Repo struct {
	db *dbconn.DB
}

/*Create new repo with acce3ss to mysql database*/
func New(conn *dbconn.DB) *Repo {
	repo := new(Repo)
	repo.db = conn
	return repo
}

/*Get user from the database by username*/
func (repo *Repo) getUser(username string, password string) (User, error) {
	
	conn := repo.db.Connection
	userStmt, err := conn.Prepare("SELECT UserId, Username FROM User WHERE username = ? AND password = ?")
	if err != nil {
		panic(err.Error())
	}
	defer userStmt.Close()
	
	var user User
	err = userStmt.QueryRow(username, password).Scan(&user)

	return user, err
}