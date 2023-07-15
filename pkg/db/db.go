package db

//--
// Database Setup
//--

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST     = "postgres-db"
	PORT     = 5432
	PASSWORD = "buck"
	DB_NAME  = "buckets_db"
)

type Database struct {
	connection *sql.DB
}

var ErrorNoMatch = fmt.Errorf("the requested record does not exist in the table")

func InitializeDB(username, password, database string) (Database, error) {
	db := Database{}

	// Data source name for the postgres driver
	dbDSN := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, username, password, database)

	conn, err := sql.Open("postgres", dbDSN)

	if err != nil {
		return db, err
	}

	db.connection = conn
	err = db.connection.Ping()

	if err != nil {
		return db, err
	}

	log.Println("Database initialized and connection established.")

	return db, nil
}
