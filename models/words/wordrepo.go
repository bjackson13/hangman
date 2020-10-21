package words

import (
	"github.com/bjackson13/hangman/models"
	"database/sql"
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

/*GetWord retireve word from DB*/
func (repo *Repo) GetWord(wordID int) (*GameWord, error) {
	wordStmt, err := repo.DB.Prepare("SELECT * FROM Words WHERE WordId = ?")
	defer wordStmt.Close()

	if err != nil {
		return nil, err
	}
	
	var word GameWord
	err = wordStmt.QueryRow(wordID).Scan(&word.WordID, &word.Word, &word.correctGuesses, &word.incorrectGuesses)
	return &word, err
}

/*AddWord add a word to a game*/
func (repo *Repo) AddWord(word string) (int, error) {
	wordStmt, err := repo.DB.Prepare("INSERT INTO Words(Word) VALUE (?)")
	defer wordStmt.Close()
	if err != nil {
		return -1, err
	}

	res, err := wordStmt.Exec(word)
	if err != nil {
		return -1, err
	}
	lastID, err := res.LastInsertId()
	return int(lastID), err
}

/*UpdateWordGuesses update the gueeses of a word*/
func (repo *Repo) UpdateWordGuesses(word GameWord) error {
	wordStmt, err := repo.DB.Prepare("UPDATE Words SET CorrectGuesses = ?, IncorrectGuesses = ? WHERE WordId = ?")
	defer wordStmt.Close()
	if err != nil {
		return err
	}

	_, err = wordStmt.Exec(&word.correctGuesses, &word.incorrectGuesses, &word.WordID)
	return err
}

/*UpdateWord update a game word and reset guesses*/
func (repo *Repo) UpdateWord(newWord string, wordID int) error {
	wordStmt, err := repo.DB.Prepare("UPDATE Words SET Word = ?, CorrectGuesses = '', IncorrectGuesses = '' WHERE WordId = ?")
	defer wordStmt.Close()
	if err != nil {
		return err
	}

	_, err = wordStmt.Exec(newWord, wordID)
	return err
}

/*DeleteWord - remove a user from the DB*/
func (repo *Repo) DeleteWord(wordID int) error {
	wordStmt, err := repo.DB.Prepare("DELETE FROM Words WHERE WordId = ?")
	defer wordStmt.Close()
	if err != nil {
		return err
	}

	_, err = wordStmt.Exec(wordID)
	return err
}