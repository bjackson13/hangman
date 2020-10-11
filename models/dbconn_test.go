package dbconn

import (
	"testing"
)

func TestConnect(t *testing.T) {
	db, err := Connect("brenbren", "password", "btj9560")
	if err != nil {
		t.Errorf("Could not connect to DB")
	}
	err = db.connection.Ping()
	if err != nil {
		t.Errorf("Could not ping DB")
	}

	err = db.connection.Close()
	if err != nil {
		t.Errorf("Failed to close DB")
	}
}