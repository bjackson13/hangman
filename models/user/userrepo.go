package user 

import (
	"github.com/bjackson13/hangman/models"
	"fmt"
)

/*
	Struct for CRUDing users from the database
*/
type Repo struct {
	db *sql.DB
}

/*Create new repo with acce3ss to mysql database*/
func (repo *Repo) New(conn *sql.DB) Repo{
	repo.db = conn
	return repo
}

/*Get user from the database by username*/
func (repo *Repo) getUser(username string, password string) User {
	
	userStmt, err := repo.db.Prepare("SELECT UserId, Username FROM User WHERE username = ? AND password = ?")
	if err != nil {
		panic(err.Error())
	}
	defer userStmt.Close()
	
	var user User
	err = userStmt.QueryRow(username, password).Scan(&user)
	if err != nil {
		panic(err.Error())
	}

	return user
}