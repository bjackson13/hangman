package lobby

import (
	"errors"
	"github.com/bjackson13/hangman/models/user"
	"github.com/bjackson13/hangman/models/lobby"
)

/*Service struct to bind our service functions to*/
type Service struct {}

func (service *Service) getLobbyUsers() ([]user.User, error) {
	lobbyRepo, err := lobby.NewRepo()
	defer lobbyRepo.Close()
	if err != nil {
		return nil, err
	}

	return lobbyRepo.GetAllLobbyUsers()
}

func (service *Service) addUser(userID int) error {
	lobbyRepo, err := lobby.NewRepo()
	defer lobbyRepo.Close()
	if err != nil {
		return err
	}
	return lobbyRepo.AddLobbyUser(userID)
}

func (service *Service) userIsInLobby(userID int) bool {
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

func (service *Service) inviteUserToPlay(invitee int, inviter int) error {
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

func (service *Service) removeUser(userID int) error {
	lobbyRepo, err := lobby.NewRepo()
	defer lobbyRepo.Close()
	if err != nil {
		return err
	}
	err = lobbyRepo.RemoveLobbyUser(userID)
	return err
}