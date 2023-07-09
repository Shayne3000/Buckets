package handler

//--
// Main router setup
//--

import (
	"net/http"

	"github.com/Shayne3000/Buckets/pkg/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var databaseInstance db.Database

// Returns the router instance
func NewRouter(db db.Database) http.Handler {
	router := chi.NewRouter()
	databaseInstance = db

	// adding custom handlers for 404 and 405 errors
	router.NotFound(notFound)
	router.MethodNotAllowed(methodNotAllowed)

	// implement a sub router items and
	// route all http requests to the /items path to the items sub router
	router.Route("/items", items)

	return router
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(404)
	render.Render(w, r, ErrorNotFound)
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrorMethodNotAllowed)
}
