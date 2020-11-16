package games

import (
	"github.com/bjackson13/hangman/models/game"
	"github.com/bjackson13/hangman/models/words"
	"errors"
	"strings"
	"strconv"
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
	defer gameRepo.Close()
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

/*RestartGame swap user roles and restart game*/
func (s *Service) RestartGame(gameID int) error {
	gameRepo, err := games.NewRepo()
	defer gameRepo.Close()
	if err != nil {
		return nil
	}

	return gameRepo.SwapUsers(gameID)
}

/*CheckPendingGuesses see if game has a guess*/
func (s *Service) CheckPendingGuesses(gameID, userID int) (string,error) {
	gameRepo, err := games.NewRepo()
	defer gameRepo.Close()
	if err != nil {
		return "", err
	}
	
	guess, err := gameRepo.GetGuess(gameID, userID) 

	if guess == "" {
		return "", errors.New("No Guess") //use an error to signal if we have a guess or not
	}

	return guess, err
}

/*MakeGuess submit a guess*/
func (s *Service) MakeGuess(gameID, userID, wordID int, guess string) GuessResult {
	wordChan := make(chan *words.GameWord)
	go func() {
		wordsRepo, err := words.NewRepo()
		defer wordsRepo.Close()
		word, err := wordsRepo.GetWord(wordID)
		if err != nil {
			wordChan <- nil 
			return
		} 
		wordChan <- word
	}()

	/*We save time (not resources unfortunatley) by making checks now and creating repos in case they are needed*/
	gameRepo, err := games.NewRepo()
	defer gameRepo.Close()
	if err != nil {
		return GuessResult{false, false, err}
	}
	
	if len(guess) <= 0 {
		return GuessResult{false, false, errors.New("No guess")}
	}

	/*
		strip guess to just first character.
		we convert to rune array first so the string is utf-8 encoded
	*/
	guess = strings.ToUpper(string([]rune(guess)[0]))

	word := <- wordChan
	if !word.GuessAlreadyMade(guess) {
		return GuessResult {
				word.IsCompleted(), 
				word.GuessLimitExceeded(), 
				gameRepo.AddGuess(guess, gameID, userID),
			}
	}

	return GuessResult{false, false, errors.New("Guess already made")}
}

/*DenyGuess deny a guess*/
func (s *Service) DenyGuess(game games.Game) GuessResult {
	gameRepo, err := games.NewRepo()
	defer gameRepo.Close()
	if err != nil {
		return GuessResult{false, false, err}
	}
	wordsRepo, err := words.NewRepo()
	defer wordsRepo.Close()
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
		guessChan <- GuessResult {
						word.IsCompleted(), 
						word.GuessLimitExceeded(), 
						wordsRepo.UpdateWordGuesses(*word),
					}
	}()

	go func() {
		gameRepo.RemoveGuess(game.GameID)
	}()

	return <- guessChan

}

/*AcceptGuess accept a guess and place the letter in the proper indexes. Indexes needs to be 0-based*/
func (s *Service) AcceptGuess(game games.Game, indexes []string) error {
	/*MAke the necessary repos*/
	gameRepo, err := games.NewRepo()
	defer gameRepo.Close()
	if err != nil {
		return err
	}
	wordsRepo, err := words.NewRepo()
	defer wordsRepo.Close()
	if err != nil {
		return err
	}

	//convert indexes strings to ints
	var idx = []int{}
	for _,v := range indexes {
		num,_ := strconv.Atoi(v)
		idx = append(idx, num)
	}

	//go submit the guess
	errChan := make(chan error)
	go func() {
		word, wErr := wordsRepo.GetWord(game.WordID)
		if wErr != nil {
			errChan <- err
		}
		word.AddCorrectGuess(game.PendingGuess, idx)
		errChan <- wordsRepo.UpdateWordGuesses(*word)
	}()

	//at the same time remove our pending guess
	go func() {
		gameRepo.RemoveGuess(game.GameID)
	}()

	return <- errChan //return result of accepting guess
}

/*AddWord add the new game word*/
func (s *Service) AddWord(gameID, wordLength int) error {
	
	if wordLength <= 0 || wordLength > 15 {
		return errors.New("Word is too long")
	} 

	gameRepo, err := games.NewRepo()
	defer gameRepo.Close()
	wordsRepo, err := words.NewRepo()
	defer wordsRepo.Close()
	if err != nil {
		return err
	}

	wordID, err := wordsRepo.AddWord(wordLength)
	if err != nil {
		return err
	}
	
	return gameRepo.UpdateWord(gameID, wordID)
}

/*GetGuesses return incorrect guesses for the game word*/
func (s *Service) GetGuesses(wordID int) ([]string, []string, error) {
	wordsRepo, err := words.NewRepo()
	defer wordsRepo.Close()
	if err != nil {
		return []string{},[]string{},err
	}

	word, err := wordsRepo.GetWord(wordID)
	if err != nil {
		return []string{},[]string{},err
	}

	/*This is totally unessesary but I wanted to try out anonymous structs */
	guesses := struct {
		Correct []string
		Incorrect []string
	}{
		Correct: word.GetCorrectGuesses(), 
		Incorrect: word.GetIncorrectGuesses(),
	}
	

	return guesses.Correct, guesses.Incorrect, nil
}

/*GetGameStatus check if game end conditions have been met*/
func (s *Service) GetGameStatus(wordID int) GuessResult {
	wordsRepo, err := words.NewRepo()
	defer wordsRepo.Close()
	if err != nil {
		return GuessResult{false, false, err}
	}

	word, err := wordsRepo.GetWord(wordID)
	if err != nil {
		return GuessResult{false, false, err}
	}

	completedChan := make(chan bool)
	exceededChan := make (chan bool)

	/*Run these checks in parallel*/
	go func() {
		completedChan <- word.IsCompleted()
	}()

	go func() {
		exceededChan <- word.GuessLimitExceeded()
	}()

	return GuessResult{ <- completedChan, <- exceededChan , nil}
}