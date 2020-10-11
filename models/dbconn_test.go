package dbconn

import (
	"testing"
	"os"
	"github.com/joho/godotenv"
)

func TestConnect(t *testing.T) {
	envErr := godotenv.Load()
	if envErr != nil {
		t.Errorf("Error loading .env file")
	}

  	mysqlUser := os.Getenv("MYSQL_TEST_USER")
	mysqlPass := os.Getenv("MYSQL_TEST_PASSWORD")
	mysqlDB := os.Getenv("MYSQL_TEST_DB")

	db, err := Connect(mysqlUser, mysqlPass, mysqlDB)
	if err != nil {
		t.Errorf("Could not connect to DB")
	}

	err = db.Connection.Ping()
	if err != nil {
		t.Errorf("Could not ping DB")
	}

	err = db.Connection.Close()
	if err != nil {
		t.Errorf("Failed to close DB")
	}
}