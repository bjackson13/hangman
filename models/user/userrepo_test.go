package user

import (
	"testing"
	"os"
	"time"
	"github.com/bjackson13/hangman/services/config"
)

var userRepo *Repo
var insertedID int
var lastLoginStamp int64

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
	
	userRepo, err = NewRepo()
	if err != nil {
		panic(err.Error())
	}
}

func teardown() {
	err := userRepo.Close()
	if err != nil {
		panic(err.Error())
	}
}

func TestNewRepo(t *testing.T) {
	err := userRepo.DB.Ping()
	if err != nil {
		t.Errorf("Failed to create UserRepo with Database connection! %s", err.Error())
	}
}

func TestAddUser(t *testing.T) {
	lastLoginStamp = time.Now().Unix()
	insertedUserID, err := userRepo.AddUser("test", "test", "192.168.1.1", "Mozilla...", lastLoginStamp)
	if err != nil {
		t.Errorf("Error inserting new user into DB, %s", err.Error())
	}
	insertedID = insertedUserID
}

func TestGetUser(t *testing.T) {
	user,err := userRepo.GetUser("test")
	if err != nil{
		t.Errorf("Error retrieving user: %s", err.Error())
	}
	if user.Username != "test" && user.UserID != insertedID && user.IP != "192.168.1.1" && user.UserAgent != "Mozilla..." && user.LastLogin != lastLoginStamp {
		t.Errorf("Failed to retrieve user, got: %v", user)
	}
}

func TestUpdateUser(t *testing.T) {
	user := NewUser("updated", "updated", "192.1.1.1", "Chrome", 0)
	user.UserID = insertedID // update our insterted user

	updatedUserCount,err := userRepo.UpdateUser(*user)
	if err != nil || updatedUserCount == 0 {
		t.Errorf("Error updating user: %s", err.Error())
	}
}

func TestUpdateUserIdentifiers(t *testing.T) {
	lastLoginStamp = time.Now().Unix()
	updatedUserCount,err := userRepo.UpdateUserIdentifiers(insertedID, "169.1.1.1", "Testing", lastLoginStamp)
	if err != nil || updatedUserCount == 0 {
		t.Errorf("Error updating user identifiers: %s", err.Error())
	}
}

func TestDeleteUser(t *testing.T) {
	deleted, err := userRepo.DeleteUser(insertedID)
	if err != nil || deleted == 0 {
		t.Errorf("Error deleting user: %s", err.Error())
	}
}