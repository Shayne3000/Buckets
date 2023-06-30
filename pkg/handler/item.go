package handler

// holds the items sub router, the items context chi middleware and the handler functions for the HTTP requests that come to the /items path

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// Used for passing the itemId URL parameter across middlewares and request handlers using Go's context.
var itemIDKey = "ItemID"

func items(router chi.Router) {
	router.Get("/", getAllItems) // handle GET requests to the root path using the getAllItems handler.
	router.Post("/", createItem)
	router.Route("/{itemId}", func(r chi.Router) {
		router.Use(ItemsCtx)
		router.Get("/", getItem)
		router.Put("/", updateItem)
		router.Delete("/" deleteItem)
	})
}

// Chi middleware function that extracts the itemId URL parameter from the request URL for use in the handlers, 
// verifies that it exists and is valid, adds it to the request context using itemIDKey (to persist the value across API boundaries), 
// and calls the next handler in the chain.
func ItemsCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemId := chi.URLParam(r, "itemId")

		if itemId == "" {
			render.Render(w, r, RenderInvalidRequestError(fmt.Errorf("item Id is required.")))
			return
		}

		id, err := strconv.Atoi(itemId)

		if err != nil {
			render.Render(w, r, RenderInvalidRequestError(fmt.Errorf("invalid item Id.")))
		}
		
		// Add the itemId given as id to the request context to hold it across API boundaries i.e. handlers?
		ctx := context.WithValue(r.Context(), itemIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
