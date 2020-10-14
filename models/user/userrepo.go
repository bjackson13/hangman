package user 

import (
	"github.com/bjackson13/hangman/models"
)

/*Repo - Struct for CRUDing users from the database*/
type Repo struct {
	db *dbconn.DB
}

/*NewRepo - Create new repo with acce3ss to mysql database*/
func NewRepo(conn *dbconn.DB) *Repo {
	repo := new(Repo)
	repo.db = conn
	return repo
}

/*GetUser - get a user from the database with provided credentials*/
func (repo *Repo) GetUser(username string) (*User, error) {
	
	conn := repo.db.Connection
	userStmt, err := conn.Prepare("SELECT UserId, Username, password, IP, UserAgent, LastLogin FROM User WHERE username = ?")
	if err != nil {
		return nil, err
	}
	defer userStmt.Close()
	
	var user User
	err = userStmt.QueryRow(username).Scan(&user.UserID, &user.Username, &user.password, &user.IP, &user.UserAgent, &user.LastLogin)

	return &user, err
}

/*AddUser - Add a user to database*/
func (repo *Repo) AddUser(username string, password string, ip string, useragent string, lastlogin int64) (int64, error) {

	conn := repo.db.Connection
	userStmt, err := conn.Prepare("INSERT INTO User(Username, Password, IP, UserAgent, LastLogin) VALUES (?,?,?,?,?)")
	if err != nil {
		return -1, err
	}
	defer userStmt.Close()

	res, err := userStmt.Exec(username, password, ip, useragent, lastlogin)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

/*UpdateUser - update an entire user. Returns rows affected or an error*/
func (repo *Repo) UpdateUser(user User) (int64, error) {
	conn := repo.db.Connection
	userStmt, err := conn.Prepare("UPDATE User SET Username = ?, Password = ?, IP = ?, UserAgent = ? WHERE UserId = ?")
	if err != nil {
		return 0,err
	}
	defer userStmt.Close()

	res, err := userStmt.Exec(user.Username, user.GetPassword(), user.IP, user.UserAgent, user.UserID)
	if err != nil {
		return 0,err
	}

	rows, err := res.RowsAffected()
	return rows, err
}

/*UpdateUserIdentifiers - update just the UserAgent and IP fields of a user*/
func (repo *Repo) UpdateUserIdentifiers(userID int64, ip string, useragent string, lastlogin int64) (int64, error) {
	conn := repo.db.Connection
	userStmt, err := conn.Prepare("UPDATE User SET IP = ?, UserAgent = ?, LastLogin = ? WHERE UserId = ?")
	if err != nil {
		return 0,err
	}
	defer userStmt.Close()

	res, err := userStmt.Exec(ip, useragent, lastlogin, userID)
	if err != nil {
		return 0,err
	}
	
	rows, err := res.RowsAffected()
	return rows, err
}

/*DeleteUser - remove a user from the DB*/
func (repo *Repo) DeleteUser(userID int64) (int64, error) {
	conn := repo.db.Connection
	userStmt, err := conn.Prepare("DELETE FROM User WHERE UserId = ?")
	if err != nil {
		return 0,err
	}
	defer userStmt.Close()

	res, err := userStmt.Exec(userID)
	if err != nil {
		return 0,err
	}
	
	rows, err := res.RowsAffected()
	return rows, err
}