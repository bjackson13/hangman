package user

import (
	"github.com/bjackson13/hangman/models"
	"testing"
	"os"
	"github.com/joho/godotenv"
)

var conn *dbconn.DB

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

func TestNew(t *testing.T) {
	userRepo := New(conn)
	err := userRepo.db.Connection.Ping()
	if err != nil {
		t.Errorf("Failed to create UserRepo with Database connection!")
	}
}

func TestGetUser(t *testing.T) {
	userRepo := New(conn)
	user,err := userRepo.GetUser("test", "test")
	if err != nil{
		t.Errorf("Error retrieving user: %s", err.Error())
	}
	if user.Username != "test" && user.UserID != 1 && user.IP != "192.168.1.1" && user.UserAgent != "Mozilla..." {
		t.Errorf("Failed to retrieve user, got: %v", user)
	}
}