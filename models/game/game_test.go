package games

import (
	"testing"
	"os"
	"github.com/bjackson13/hangman/services/config"
)

var gameRepo *Repo
var gameID int
var wordID int =  1 //test ID
var guessingUser int = 2
var creatingUser int = 3

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
	
	gameRepo, err = NewRepo()
	if err != nil {
		panic(err.Error())
	}
}

func teardown() {
	err := gameRepo.Close()
	if err != nil {
		panic(err.Error())
	}
}

func TestNewRepo(t *testing.T) {
	err := gameRepo.DB.Ping()
	if err != nil {
		t.Errorf("Failed to create gameRepo with Database connection! %s", err.Error())
	}
}

func TestAddGame(t *testing.T) {
	id, err := gameRepo.AddGame(guessingUser,creatingUser)
	if err != nil {
		t.Errorf("Error creating new game: %v", err.Error())
	}

	if id <= 0 {
		t.Errorf("Error creating new game: %v", id)
	}
	gameID = id
}

func TestUpdateWord(t *testing.T) {
	err := gameRepo.UpdateWord(gameID, wordID)
	if err != nil {
		t.Errorf("Error updating word in game: %v", err.Error())
	}
}

func TestGetGameByUser(t *testing.T) {
	game,err := gameRepo.GetGameByUser(guessingUser)
	if err != nil {
		t.Errorf("Error getting game by user id: %v", err.Error())
	}

	if game.GameID != gameID || game.WordID != wordID || game.GuessingUserID != guessingUser  {
		t.Errorf("Error getting game by user id: %v", game)
	}
	t.Log(game)
}

func TestAddGuess(t *testing.T) {
	err := gameRepo.AddGuess("h", gameID, guessingUser)
	if err != nil {
		t.Errorf("Error adding guess to game: %v", err.Error())
	}
}

func TestGetGuess(t *testing.T) {
	guess, err := gameRepo.GetGuess(gameID, creatingUser)
	if err != nil {
		t.Errorf("Error getting guess: %v", err.Error())
	}

	if guess != "h" {
		t.Errorf("Wrong value returned for guess: %v", guess)
	}
}

func TestRemoveGuess(t *testing.T) {
	err := gameRepo.RemoveGuess(gameID)
	if err != nil {
		t.Errorf("Error removing guess: %v", err.Error())
	}
}

func TestSwapUsers(t *testing.T) {
	err := gameRepo.SwapUsers(gameID)
	if err != nil {
		t.Errorf("Error swapping game users: %v", err.Error())
	}
}

func TestSwappedUsersConfirm(t *testing.T) {
	game,err := gameRepo.GetGameByUser(guessingUser)
	if err != nil {
		t.Errorf("Error getting game by user id after swap: %v", err.Error())
	}

	if game.GameID != gameID || game.WordID != -1 || game.GuessingUserID != creatingUser || game.WordCreatorID != guessingUser {
		t.Errorf("Error swapping users for game: %v", game)
	}
}

func TestRemoveGame(t *testing.T) {
	err := gameRepo.RemoveGame(gameID)
	if err != nil {
		t.Errorf("Error removing game: %v", err.Error())
	}
}

