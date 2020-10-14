package dbconn

import (
	"testing"
)

func TestConnect(t *testing.T) {
	db, err := Connect()
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