package words

import (
	"github.com/bjackson13/hangman/models"
	"database/sql"
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

/*GetWord retireve word from DB*/
func (repo *Repo) GetWord(wordID int) (*GameWord, error) {
	wordStmt, err := repo.DB.Prepare("SELECT * FROM Words WHERE WordId = ?")
	defer wordStmt.Close()

	if err != nil {
		return nil, err
	}
	
	var word GameWord
	var correctGuesses sql.NullString
	var incorrectGuesses sql.NullString
	err = wordStmt.QueryRow(wordID).Scan(&word.WordID, &word.Length, &correctGuesses, &incorrectGuesses)
	if correctGuesses.Valid {
		word.SetCorrectGuesses(correctGuesses.String)
	} else {
		word.SetCorrectGuesses("")
	}
		
	if incorrectGuesses.Valid {
		word.SetIncorrectGuesses(incorrectGuesses.String)
	} else {
		word.SetIncorrectGuesses("")
	}
	return &word, err
}

/*AddWord add a word to a game*/
func (repo *Repo) AddWord(length int) (int, error) {
	wordStmt, err := repo.DB.Prepare("INSERT INTO Words(WordLength) VALUE (?)")
	defer wordStmt.Close()
	if err != nil {
		return -1, err
	}

	res, err := wordStmt.Exec(length)
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

	_, err = wordStmt.Exec(word.CorrectToString(), word.IncorrectToString(), &word.WordID)
	return err
}

/*UpdateWord update a game word and reset guesses*/
func (repo *Repo) UpdateWord(newLength int, wordID int) error {
	wordStmt, err := repo.DB.Prepare("UPDATE Words SET WordLength = ?, CorrectGuesses = '', IncorrectGuesses = '' WHERE WordId = ?")
	defer wordStmt.Close()
	if err != nil {
		return err
	}

	_, err = wordStmt.Exec(newLength, wordID)
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