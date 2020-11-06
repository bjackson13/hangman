package games

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

/*AddGame create a new game in the database for 2 users*/
func (repo *Repo) AddGame(guessingUserID, wordCreatorID int) (int, error) {
	gameStmt, err := repo.DB.Prepare("INSERT INTO Games (GuessingUserId, WordCreatorId) VALUES (?,?)")
	defer gameStmt.Close()
	if err != nil {
		return -1, err
	}

	res, err := gameStmt.Exec(guessingUserID, wordCreatorID)
	if err != nil {
		return -1, err
	}
	lastID, err := res.LastInsertId()
	return int(lastID), err
}

/*GetGameByUser get a game by querying the guesser or word creator*/
func (repo *Repo) GetGameByUser(userID int) (*Game, error) {
	gameStmt, err := repo.DB.Prepare("SELECT * FROM Games WHERE GuessingUserId = ? OR WordCreatorId = ?")
	defer gameStmt.Close()
	if err != nil {
		return nil, err
	}

	var wordID sql.NullInt32
	var guess sql.NullString
	var game Game
	err = gameStmt.QueryRow(userID, userID).Scan(&game.GameID, &wordID, &game.GuessingUserID, &game.WordCreatorID, &guess)
	
	if err != nil {
		return nil, err
	}
	
	//word ID
	if !wordID.Valid {
		game.WordID = -1
	} else {
		game.WordID = int(wordID.Int32)
	}

	//Pending guess
	if !guess.Valid {
		game.PendingGuess = ""
	} else {
		game.PendingGuess = guess.String
	}

	return &game, err
}

/*UpdateWord updates the WordId column of a game*/
func (repo *Repo) UpdateWord(gameID, wordID int) error {
	gameStmt, err := repo.DB.Prepare("UPDATE Games SET WordId = ? WHERE GameId = ?")
	defer gameStmt.Close()
	if err != nil {
		return err
	}

	_, err = gameStmt.Exec(wordID, gameID)
	return err
}

/*AddGuess Update the PendingGuess column of a game with the guessing users guess. Will fail if anyone but guessing user makes guess*/
func (repo *Repo) AddGuess(guess string, gameID int, userID int) error {
	gameStmt, err := repo.DB.Prepare("UPDATE Games SET PendingGuess = ? WHERE GameId = ? AND GuessingUserId = ?")
	defer gameStmt.Close()
	if err != nil {
		return err
	}

	_, err = gameStmt.Exec(guess, gameID, userID)
	return err
}

/*GetGuess check if a game has a guess and return the guess*/
func (repo *Repo) GetGuess(gameID, userID int) (string,error) {
	gameStmt, err := repo.DB.Prepare("SELECT PendingGuess FROM Games WHERE GameId = ? AND WordCreatorId = ?")
	defer gameStmt.Close()
	if err != nil {
		return "", err
	}

	var guess sql.NullString
	err = gameStmt.QueryRow(gameID, userID).Scan(&guess)
	if !guess.Valid {
		return "", err
	}
	return guess.String, err
}

/*RemoveGuess set PendingGuess to NULL*/
func (repo *Repo) RemoveGuess(gameID int) error {
	gameStmt, err := repo.DB.Prepare("UPDATE Games SET PendingGuess = NULL WHERE GameId = ?")
	defer gameStmt.Close()
	if err != nil {
		return err
	}

	_, err = gameStmt.Exec(gameID)
	return err
}

/*SwapUsers make the guessing user the word creator and the word creator the guessing user*/
func (repo *Repo) SwapUsers(gameID int) error {
	gameStmt, err := repo.DB.Prepare("UPDATE Games SET GuessingUserId = (@tmp:=GuessingUserId), GuessingUserId = WordCreatorId, WordCreatorId = @tmp, PendingGuess = NULL WHERE GameId = ?")
	defer gameStmt.Close()
	if err != nil {
		return err
	}

	_, err = gameStmt.Exec(gameID)
	return err
} 

/*RemoveGame remove game from Games table by ID. All other game related tables should be cleared before this*/
func (repo *Repo) RemoveGame(gameID int) error {
	gameStmt, err := repo.DB.Prepare("DELETE FROM Games WHERE GameId = ?")
	defer gameStmt.Close()
	if err != nil {
		return err
	}

	_, err = gameStmt.Exec(gameID)
	return err
} 