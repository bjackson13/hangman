package games

import (
	"github.com/bjackson13/hangman/models/game"
	"github.com/bjackson13/hangman/models/words"
	"errors"
)

/*Service struct to bind our service functions to*/
type Service struct {}

/*GuessResult tells us the outcome of a game based on a guess*/
type GuessResult struct {
	WordComplete bool
	LimitExceeded bool
	Error error
}

/*NewService produce a new service*/
func NewService() *Service {
	return new(Service)
}

/*GetUserGame get the game a user is currently in. Nil if none*/
func (s *Service) GetUserGame(userID int) *games.Game {
	gameRepo, err := games.NewRepo()
	if err != nil {
		return nil
	}

	userGame, err := gameRepo.GetGameByUser(userID)
	if err != nil {
		return nil
	}

	return userGame
}

/*EndGame end a game by removing it*/
func (s *Service) EndGame(gameID int) error {
	gameRepo, err := games.NewRepo()
	if err != nil {
		return err
	}

	return gameRepo.RemoveGame(gameID)
	 
}

/*CheckGuesses see if game has a guess*/
func (s *Service) CheckGuesses(gameID, userID int) (string,error) {
	gameRepo, err := games.NewRepo()
	if err != nil {
		return "", err
	}
	
	guess, err := gameRepo.GetGuess(gameID, userID) 

	if guess == "" {
		return "", errors.New("No Guess")
	}

	return guess, err
}

/*MakeGuess submit a guess*/
func (s *Service) MakeGuess(gameID, userID, wordID int, guess string) error {
	guessChan := make(chan bool)
	go func() {
		wordsRepo, err := words.NewRepo()
		defer wordsRepo.Close()
		word, err := wordsRepo.GetWord(wordID)
		if err != nil {
			guessChan <- true // use true to skip adding guess if we  can't check
			return
		} 
		guessChan <- word.GuessAlreadyMade(guess)
	}()
	
	guessMadePreviously := <- guessChan

	/*We save time (not resources unfortunatley) by making checks now and creating repos in case they are needed*/
	gameRepo, err := games.NewRepo()
	if err != nil {
		return err
	}
	
	if len(guess) <= 0 {
		return errors.New("No guess")
	}

	/*
		strip guess to just first character.
		we convert to rune array first so the string is utf-8 encoded
	*/
	guess = string([]rune(guess)[0])

	if !guessMadePreviously {
		return gameRepo.AddGuess(guess, gameID, userID)

	}

	return errors.New("Guess already made")
}

/*DenyGuess deny a guess*/
func (s *Service) DenyGuess(game games.Game) GuessResult {
	gameRepo, err := games.NewRepo()
	wordsRepo, err := words.NewRepo()
	if err != nil {
		return GuessResult{false, false, err}
	}

	guessChan := make(chan GuessResult)
	go func() {
		word, wErr := wordsRepo.GetWord(game.WordID)
		if wErr != nil {
			guessChan <- GuessResult{false, false, wErr}
		}
		word.AddIncorrectGuess(game.PendingGuess)
		guessChan <- GuessResult{word.IsCompleted(), word.GuessLimitExceeded(), wordsRepo.UpdateWordGuesses(*word)}
	}()

	go func() {
		gameRepo.RemoveGuess(game.GameID)
	}()

	return <- guessChan

}

/*AddWord add the new game word*/
func (s *Service) AddWord(gameID, wordLength int) error {
	
	if wordLength <= 0 || wordLength > 15 {
		return errors.New("Word is too long")
	} 

	gameRepo, err := games.NewRepo()
	wordsRepo, err := words.NewRepo()
	if err != nil {
		return err
	}

	wordID, err := wordsRepo.AddWord(wordLength)
	if err != nil {
		return err
	}
	
	return gameRepo.UpdateWord(gameID, wordID)
}