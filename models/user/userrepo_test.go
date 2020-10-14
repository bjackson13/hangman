package user

import (
	"github.com/bjackson13/hangman/models"
	"testing"
	"os"
	"time"
)

var conn *dbconn.DB
var insertedID int64
var lastLoginStamp int64

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	db, err := dbconn.Connect()
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

func TestNewRepo(t *testing.T) {
	userRepo := NewRepo(conn)
	err := userRepo.db.Connection.Ping()
	if err != nil {
		t.Errorf("Failed to create UserRepo with Database connection! %s", err.Error())
	}
}

func TestAddUser(t *testing.T) {
	userRepo := NewRepo(conn)

	lastLoginStamp = time.Now().Unix()
	insertedUserID, err := userRepo.AddUser("test", "test", "192.168.1.1", "Mozilla...", lastLoginStamp)
	if err != nil {
		t.Errorf("Error inserting new user into DB, %s", err.Error())
	}
	insertedID = insertedUserID
}

func TestGetUser(t *testing.T) {
	userRepo := NewRepo(conn)
	user,err := userRepo.GetUser("test")
	if err != nil{
		t.Errorf("Error retrieving user: %s", err.Error())
	}
	if user.Username != "test" && user.UserID != insertedID && user.IP != "192.168.1.1" && user.UserAgent != "Mozilla..." && user.LastLogin != lastLoginStamp {
		t.Errorf("Failed to retrieve user, got: %v", user)
	}
}

func TestUpdateUser(t *testing.T) {
	userRepo := NewRepo(conn)
	user := NewUser("updated", "updated", "192.1.1.1", "Chrome", 0)
	user.UserID = insertedID // update our insterted user

	updatedUserCount,err := userRepo.UpdateUser(*user)
	if err != nil || updatedUserCount == 0 {
		t.Errorf("Error updating user: %s", err.Error())
	}
}

func TestUpdateUserIdentifiers(t *testing.T) {
	userRepo := NewRepo(conn)
	lastLoginStamp = time.Now().Unix()
	updatedUserCount,err := userRepo.UpdateUserIdentifiers(insertedID, "169.1.1.1", "Testing", lastLoginStamp)
	if err != nil || updatedUserCount == 0 {
		t.Errorf("Error updating user identifiers: %s", err.Error())
	}
}

func TestDeleteUser(t *testing.T) {
	userRepo := NewRepo(conn)

	deleted, err := userRepo.DeleteUser(insertedID)
	if err != nil || deleted == 0 {
		t.Errorf("Error deleting user: %s", err.Error())
	}
}