package games

import (
	"github.com/bjackson13/hangman/models/game"
	"errors"
)

/*Service struct to bind our service functions to*/
type Service struct {}

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
func (s *Service) MakeGuess(gameID, userID int, guess string) error {
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

	return gameRepo.AddGuess(guess, gameID, userID)
	 
}

/*DenyGuess deny a guess*/
func (s *Service) DenyGuess(gameID int) error {
	gameRepo, err := games.NewRepo()
	if err != nil {
		return err
	}
	
	return gameRepo.RemoveGuess(gameID)
}