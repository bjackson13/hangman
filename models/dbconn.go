package dbconn

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" /*MYSQL Driver*/
)

/* Conenct to MYSQL database */
func Connect(user string, password string, dbName string) (*sql.DB, error) {
	connString :=  user + ":" + password + "@/" + dbName
	DB, err := sql.Open("mysql", connString)
	return DB, err
}