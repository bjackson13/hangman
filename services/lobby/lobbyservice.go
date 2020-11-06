package lobby

import (
	"errors"
	"github.com/bjackson13/hangman/models/user"
	"github.com/bjackson13/hangman/models/lobby"
	"github.com/bjackson13/hangman/models/game"
)

/*Service struct to bind our service functions to*/
type Service struct {}

/*NewService produce a new service*/
func NewService() *Service {
	return new(Service)
}

func (service *Service) GetLobbyUsers() ([]user.User, error) {
	lobbyRepo, err := lobby.NewRepo()
	defer lobbyRepo.Close()
	if err != nil {
		return nil, err
	}

	return lobbyRepo.GetAllLobbyUsers()
}

func (service *Service) AddUser(userID int) error {
	lobbyRepo, err := lobby.NewRepo()
	defer lobbyRepo.Close()
	if err != nil {
		return err
	}
	return lobbyRepo.AddLobbyUser(userID)
}

func (service *Service) UserIsInLobby(userID int) bool {
	lobbyRepo, err := lobby.NewRepo()
	defer lobbyRepo.Close()
	if err != nil {
		return false
	}
	inLobby, err := lobbyRepo.UserIsInLobby(userID)
	if err != nil {
		return false
	}
	return inLobby
}

func (service *Service) InviteUserToPlay(invitee, inviter int) error {
	lobbyRepo, err := lobby.NewRepo()
	defer lobbyRepo.Close()
	if err != nil {
		return err
	}

	inLobby, err := lobbyRepo.UserIsInLobby(invitee)
	if !inLobby || err != nil {
		return errors.New("Could not find user in lobby")
	}
	err = lobbyRepo.InviteUser(invitee, inviter)
	return err
}

func (service *Service) CheckInvites(userID int) (*string, int, error) {
	lobbyRepo, err := lobby.NewRepo()
	defer lobbyRepo.Close()
	if err != nil {
		return nil, -1, err
	}
	return lobbyRepo.CheckInvites(userID)
}

func (service *Service) AcceptInvite(inviteeID, inviterID int) (int, error) {
	
	//accept invite and create game
	gameRepo, err := games.NewRepo()
	defer gameRepo.Close()
	if err != nil {
		return -1, err
	}

	gameID, err := gameRepo.AddGame(inviteeID, inviterID)
	
	if err == nil {
		//after game assigned remove users from lobby
		go func () {
			lobbyRepo, _ := lobby.NewRepo()
			defer lobbyRepo.Close()
			lobbyRepo.RemoveLobbyUser(inviterID) 
			lobbyRepo.RemoveLobbyUser(inviteeID) 
		}()
	}
	
	return gameID, err
}

func (service *Service) DenyInvite(userID int) error {
	lobbyRepo, err := lobby.NewRepo()
	defer lobbyRepo.Close()
	if err != nil {
		return err
	}
	return lobbyRepo.RevokeInvite(userID)
}

func (service *Service) RemoveUser(userID int) error {
	lobbyRepo, err := lobby.NewRepo()
	defer lobbyRepo.Close()
	if err != nil {
		return err
	}
	err = lobbyRepo.RemoveLobbyUser(userID)
	return err
}