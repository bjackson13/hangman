package lobby

import (
	"testing"
	"os"
	"github.com/bjackson13/hangman/services/config"
)

var lobbyRepo *Repo
var chatID int = 1

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	var err error
	err = config.LoadEnvVariables()
	if err != nil{
		panic("Failed to load env variables")
	}
	
	lobbyRepo, err = NewRepo()
	if err != nil {
		panic(err.Error())
	}
}

func teardown() {
	err := lobbyRepo.Close()
	if err != nil {
		panic(err.Error())
	}
}

func TestNewRepo(t *testing.T) {
	err := lobbyRepo.DB.Ping()
	if err != nil {
		t.Errorf("Failed to create lobbyRepo with Database connection! %s", err.Error())
	}
}

func TestAddLobbyUser(t *testing.T) {
	err := lobbyRepo.AddLobbyUser(2) //test user
	if err != nil {
		t.Errorf("Error inserting new lobby user into DB, %s", err.Error())
	}
}

func TestGetAllLobbyUsers(t *testing.T) {
	users, err := lobbyRepo.GetAllLobbyUsers()
	if err != nil {
		t.Errorf("Error getting all lobby users from DB, %s", err.Error())
	}
	
	if len(users) <= 0 || users[0].UserID != 2 {
		t.Errorf("Did not get proper  users back, length: %v", users)
	}
}

func TestUserIsInLobby(t *testing.T) {
	isInLobby, err := lobbyRepo.UserIsInLobby(2) //test users
	if err != nil {
		t.Errorf("Error getting lobby user from DB, %s", err.Error())
	}

	if !isInLobby {
		t.Errorf("User is not in lobby")
	}
}

func TestInviteUser(t *testing.T) {
	err := lobbyRepo.InviteUser(2, 3) //test users
	if err != nil {
		t.Errorf("Error adding invite to user in lobby, %s", err.Error())
	}
}

func TestCheckInvites(t *testing.T) {
	inviter, id, err := lobbyRepo.CheckInvites(2) //test users
	if err != nil {
		t.Errorf("Error checking invite to user in lobby, %s", err.Error())
	}

	if *inviter != "postman" && id != 3 {
		t.Errorf("Wrong or no inviterId assigned to user 2: %v", inviter)
	}
}

func TestRevokeInvite(t *testing.T) {
	err := lobbyRepo.RevokeInvite(2) //test users
	if err != nil {
		t.Errorf("Error revoking invite from lobby user from DB, %s", err.Error())
	}
}

func TestCheckInvitesNegativeCondition(t *testing.T) {
	inviter, id, err := lobbyRepo.CheckInvites(2) //test users
	if err != nil {
		t.Errorf("Error checking invite to user in lobby, %s", err.Error())
	}

	if inviter != nil && id != -1 {
		t.Errorf("Wrong inviterId assigned to user 2: %v", inviter)
	}
}

func TestRemoveLobbyUser(t *testing.T) {
	err := lobbyRepo.RemoveLobbyUser(2) //test users
	if err != nil {
		t.Errorf("Error removing lobby user from DB, %s", err.Error())
	}
}