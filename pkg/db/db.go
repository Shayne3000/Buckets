package db

//--
// Database Setup
//--

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	HOST = "postgres-db"
	PORT = 5432
)

type Database struct {
	Connection *sql.DB
}

var ErrorNoMatch = fmt.Errorf("the requested record does not exist in the table")

func InitializeDB(username string) (Database, error) {
	db := Database{}

	// Connection string to connect to the DB
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, username, os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	conn, err := sql.Open("postgres", connectionString)

	if err != nil {
		return db, err
	}

	db.Connection = conn
	err = db.Connection.Ping()

	if err != nil {
		return db, err
	}

	log.Println("Database initialized and connection established.")

	return db, nil
}
