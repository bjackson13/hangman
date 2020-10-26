package games

import (
	"github.com/bjackson13/hangman/models"
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
func (repo *Repo) AddGame(guessingUserID int, wordCreatorID int) (int, error) {
	/*	
		Since we can't autoincrement 2 columns with MySQL: 
		we set the ChatId to a random number between 0-100000 based on the 2 users in the game.
		This has the added benefit of preventing 2 users from accidently having 2 games with each other simultaniously
	*/
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

func (repo *Repo) GetGameByUser(userID int) (*Game, error) {
	return nil, nil
}

func (repo *Repo) UpdateWord(gameID int, wordID int) error {
	return nil
}

func (repo *Repo) AddGuess(guess string, gameID int, userID int) error {
	return nil
}

func (repo *Repo) GetGuess(gameID int, userID int) (string,error) {
	return "", nil
}

func (repo *Repo) RemoveGuess(gameID int) error {
	return nil
}

func (repo *Repo) SwapUsers(gameID int) error {
	//gameStmt, err := repo.DB.Prepare("UPDATE Games SET GuessingUserId = (@tmp:=GuessingUserId),GuessingUserId = WordCreatorId, WordCreatorId = @tmp, PendingGuess = NULL WHERE GameId = ?")
	return nil
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