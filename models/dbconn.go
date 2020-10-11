package dbconn

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" /*MYSQL Driver*/
)

type DB struct {
	Connection *sql.DB
}

/* Conenct to MYSQL database */
func Connect(user string, password string, dbName string) (*DB, error) {
	connString :=  user + ":" + password + "@/" + dbName
	conn, err := sql.Open("mysql", connString)
	return &DB{conn}, err
}