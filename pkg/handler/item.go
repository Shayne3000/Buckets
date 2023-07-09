package handler

//--
// File responsible for handling CRUD HTTP requests to the item resource in the DB through the /items path.
// It holds the items sub router, the items context http middleware, and the handler functions for the HTTP requests that come to the /items path
//--

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// Key used for passing the itemId URL parameter across API boundaries/.middlewares and request handlers using Go's context.
var itemIDKey = "ItemID"

// Sub-router that handles all HTTP requests routed to the "/items" path.
func items(router chi.Router) {
	// TODO Add paginate to the GetAll request.
	router.Get("/", getAllItems) // read as tell the router to handle GET requests to the root path using the getAllItems handler.
	router.Post("/", createItem)
	router.Route("/{itemId}", func(r chi.Router) {
		router.Use(ItemsCtx)
		router.Get("/", getItem)
		router.Put("/", updateItem)
		router.Delete("/" deleteItem)
	})
}

// net/http middleware function that extracts the itemId URL parameter from the request URL 
// and saves it in the request context for use across API boundaries i.e. in several handlers.
func ItemsCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fetch the parameter at runtime
		itemId := chi.URLParam(r, "itemId")

		// verifies that the id exists
		if itemId == "" {
			render.Render(w, r, RenderInvalidRequestError(fmt.Errorf("item Id is required.")))
			return
		}

		// convert the itemId from string to int
		id, err := strconv.Atoi(itemId)

		// verify that the id is valid
		if err != nil {
			render.Render(w, r, RenderInvalidRequestError(fmt.Errorf("invalid item Id.")))
		}
		
		// Add the itemId to the request context using the itemIDKey to persist it across API boundaries
		ctx := context.WithValue(r.Context(), itemIDKey, id)

		// call the next handler in the chain
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
