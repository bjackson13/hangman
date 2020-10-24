package lobby

import (
	"github.com/bjackson13/hangman/models"
)

/*Repo - Struct for CRUDing users from the database*/
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

func (repo *Repo) AddLobbyUser(userId int) error {
	return nil
}

func (repo *Repo) GetAllLobbyUsers() ([]int, error) {
	return nil, nil
}

func (repo *Repo) UserIsInLobby(userId int) error {
	return nil
}

func (repo *Repo) RemoveLobbyUser(userId int) error {
	return nil
}






