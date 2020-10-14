package dbconn

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" /*MYSQL Driver*/
	"github.com/joho/godotenv"
	"os"
)

type DB struct {
	Connection *sql.DB
}

/*Connect to MYSQL database */
func Connect() (*DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

  	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPass := os.Getenv("MYSQL_PASSWORD")
	mysqlDB := os.Getenv("MYSQL_DB")

	connString :=  mysqlUser + ":" + mysqlPass + "@/" + mysqlDB
	conn, err := sql.Open("mysql", connString)
	return &DB{conn}, err
}