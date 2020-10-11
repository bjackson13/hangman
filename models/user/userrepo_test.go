package user

import (
	"github.com/bjackson13/hangman/models"
	"testing"
	"os"
)

var conn *dbconn.DB

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	db, err := dbconn.Connect("brenbren", "password", "btj9560")
	if err != nil {
		panic(err.Error())
	}
	conn = db
}

func teardown() {
	err := conn.Connection.Close()
	if err != nil {
		panic(err.Error())
	}
}

func TestNew(t *testing.T) {
	userRepo := New(conn)
	err := userRepo.db.Connection.Ping()
	if err != nil {
		t.Errorf("Failed to create UserRepo with Database connection!")
	}
}

func TestGetUser(t *testing.T) {
	userRepo := New(conn)
	user,err := userRepo.getUser("test", "test")
	if err != nil{
		t.Errorf("Error retrieving user: %s", err.Error())
	}
	if user.Username != "test" && user.UserId != 1 {
		t.Errorf("Failed to retrieve user, got: %v", user)
	}
}