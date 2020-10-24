package lobby

import (
	"testing"
	"os"
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
	_, err := lobbyRepo.GetAllLobbyUsers()
	if err != nil {
		t.Errorf("Error getting all lobby users from DB, %s", err.Error())
	}
}

func TestUserIsInLobby(t *testing.T) {
	err := lobbyRepo.UserIsInLobby(2) //test users
	if err != nil {
		t.Errorf("Error getting lobby user from DB, %s", err.Error())
	}
}

func TestRemoveLobbyUser(t *testing.T) {
	err := lobbyRepo.RemoveLobbyUser(2) //test users
	if err != nil {
		t.Errorf("Error removing lobby user from DB, %s", err.Error())
	}
}