package dbconn

import (
	"testing"
	"github.com/bjackson13/hangman/services/config"
)

func TestConnect(t *testing.T) {
	err := config.LoadEnvVariables()
	if err != nil{
		t.Errorf("Failed to load env variables")
	}
	
	db, err := Connect()
	if err != nil {
		t.Errorf("Could not connect to DB")
	}

	err = db.Ping()
	if err != nil {
		t.Errorf("Could not ping DB")
	}

	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close DB")
	}
}