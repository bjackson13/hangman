package games

import (
	"github.com/bjackson13/hangman/models/game"
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