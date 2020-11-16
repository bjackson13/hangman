package dbconn

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" /*MYSQL Driver*/
	"github.com/joho/godotenv"
	"os"
)

/*Repo - Struct for CRUDing from the database*/
type Repo struct {
	DB *sql.DB
}

/*Close closes the database connection*/
func (repo *Repo) Close() error {
	return repo.DB.Close()
	
}

/*Connect to MYSQL database */
func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

  	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPass := os.Getenv("MYSQL_PASSWORD")
	mysqlDB := os.Getenv("MYSQL_DB")

	connString :=  mysqlUser + ":" + mysqlPass + "@/" + mysqlDB
	conn, err := sql.Open("mysql", connString)
	return conn, err
}