package lobby

import (
	"errors"
	"github.com/bjackson13/hangman/models/user"
	"github.com/bjackson13/hangman/models/lobby"
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

func (service *Service) InviteUserToPlay(invitee int, inviter int) error {
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

func (service *Service) RemoveUser(userID int) error {
	lobbyRepo, err := lobby.NewRepo()
	defer lobbyRepo.Close()
	if err != nil {
		return err
	}
	err = lobbyRepo.RemoveLobbyUser(userID)
	return err
}