package user

import (
	"github.com/bjackson13/hangman/models"
	"testing"
	"os"
	"github.com/joho/godotenv"
)

var conn *dbconn.DB
var insertedID int64

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	/*Load in .env file: 
		 	if you want to run these tests you need
			to place a .env file in this directory
			with the same variables listed below*/
	envErr := godotenv.Load()
	if envErr != nil {
		panic(envErr.Error())
	}

  	mysqlUser := os.Getenv("MYSQL_TEST_USER")
	mysqlPass := os.Getenv("MYSQL_TEST_PASSWORD")
	mysqlDB := os.Getenv("MYSQL_TEST_DB")

	db, err := dbconn.Connect(mysqlUser, mysqlPass, mysqlDB)
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

	insertedUserID, err := userRepo.AddUser("test", "test", "192.168.1.1", "Mozilla...")
	if err != nil {
		t.Errorf("Error inserting new user into DB, %s", err.Error())
	}
	insertedID = insertedUserID
}

func TestGetUser(t *testing.T) {
	userRepo := NewRepo(conn)
	user,err := userRepo.GetUser("test", "test")
	if err != nil{
		t.Errorf("Error retrieving user: %s", err.Error())
	}
	if user.Username != "test" && user.UserID != insertedID && user.IP != "192.168.1.1" && user.UserAgent != "Mozilla..." {
		t.Errorf("Failed to retrieve user, got: %v", user)
	}
}

func TestUpdateUser(t *testing.T) {
	userRepo := NewRepo(conn)
	user := NewUser("updated", "updated", "192.1.1.1", "Chrome")
	user.UserID = insertedID // update our insterted user

	updatedUserCount,err := userRepo.UpdateUser(*user)
	if err != nil || updatedUserCount == 0 {
		t.Errorf("Error updating user: %s", err.Error())
	}
}

func TestUpdateUserIdentifiers(t *testing.T) {
	userRepo := NewRepo(conn)

	updatedUserCount,err := userRepo.UpdateUserIdentifiers(insertedID, "169.1.1.1", "Testing")
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