package main

import (
	"log"
	"net/http"

	"github.com/Shayne3000/Buckets/pkg/db"
	"github.com/Shayne3000/Buckets/pkg/handler"
)

func main() {
	serverPort := "8080"
	dbUser := "postgres"

	// setup the DB
	database, err := db.InitializeDB(dbUser)

	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer database.Connection.Close()

	// setup the router and connect it to the database instance
	router := handler.NewRouter(database)

	// start the server
	log.Fatal(http.ListenAndServe(serverPort, router))
}
