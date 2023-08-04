package handler

//--
// File responsible for handling HTTP requests to the /items path to perform CRUD ops on the item resource in the DB.
// It holds the items sub router, the items context http middleware, and the handler functions for the HTTP requests that come to the /items path
//--

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Shayne3000/Buckets/pkg/db"
	"github.com/Shayne3000/Buckets/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// A custom string type called itemID used as the type for the key used to store itemId in the request context.
// See this: https://forum.golangbridge.org/t/the-way-to-set-a-context-key-and-value/16311
type itemID string

// Key used to store and pass the itemId URL parameter across API boundaries/request handlers using Go's context.
var itemIDKey itemID = "ItemID"

// Sub-router that handles all HTTP requests routed to the "/items" path.
func items(router chi.Router) {
	// TODO Add paginate to the GetAll request.
	router.Get("/", getAllItems) // read as tell the router to handle GET requests to the root /items/ path using the getAllItems handler.
	router.Post("/", createItem)
	// route all requests with the itemId URL parameter to the sub-router
	router.Route("/{itemId}", func(r chi.Router) {
		r.Use(ItemsCtx)
		r.Get("/", getItem)
		r.Put("/", updateItem)
		r.Delete("/", deleteItem)
	})
}

// net/http middleware function that extracts the itemId URL parameter from the request URL
// and saves it in the request context for use across API boundaries i.e. in several handlers.
func ItemsCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fetch the URL parameter at runtime
		itemId := chi.URLParam(r, "itemId")

		// verifies that the id exists
		if itemId == "" {
			render.Render(w, r, RenderInvalidRequestError(fmt.Errorf("item id is required")))
			return
		}

		// convert the itemId from string to int
		id, err := strconv.Atoi(itemId)

		// verify that the id is valid
		if err != nil {
			render.Render(w, r, RenderInvalidRequestError(fmt.Errorf("invalid item id")))
		}

		// Add the itemId to the request context using the itemIDKey so as to persist and use it across API boundaries
		ctx := context.WithValue(r.Context(), itemIDKey, id)

		// call the next handler in the chain
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Create a bucket list item
func createItem(w http.ResponseWriter, r *http.Request) {
	// instance of a pointer to the item struct
	item := &models.Item{}

	// Use render.Bind to decode/unmarshall the request body into an Item model for insertion into the DB.
	if err := render.Bind(r, item); err != nil {
		render.Render(w, r, RenderInvalidRequestError(err))
		return
	}

	// Insert into the DB
	if err := databaseInstance.AddItem(item); err != nil {
		render.Render(w, r, RenderServerError(err))
		return
	}

	// Return the created item to tell the user that the request was successful
	if err := render.Render(w, r, item); err != nil {
		render.Render(w, r, RenderServerError(err))
	}
}

// Get all bucket list items
func getAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := databaseInstance.GetAllItems()

	if err != nil {
		render.Render(w, r, RenderServerError(err))
		return
	}

	// Return the requested bucket list items if the request was successful.
	if err := render.Render(w, r, items); err != nil {
		render.Render(w, r, RenderServerError(err))
	}
}

// Get a bucket list item given its id
func getItem(w http.ResponseWriter, r *http.Request) {
	// Retrieve the item ID URL parameter that was stored in the request context by the ItemsContext middleware.
	itemId := r.Context().Value(itemIDKey).(int)

	// Retrieve item from the DB given its id
	item, err := databaseInstance.GetItemById(itemId)

	if err != nil {
		if err == db.ErrorNoMatch {
			render.Render(w, r, ErrorNotFound)
		} else {
			render.Render(w, r, RenderInvalidRequestError(err))
		}
		return
	}

	// Return the requested item to the user
	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, RenderServerError(err))
	}
}

// Update an existing bucket list item given its id
func updateItem(w http.ResponseWriter, r *http.Request) {
	// Get the item id URL parameter from the request context
	itemId := r.Context().Value(itemIDKey).(int)

	updatedItemData := models.Item{}

	// Use render.Bind to decode the request body into the Item model, updatedItemData.
	if err := render.Bind(r, &updatedItemData); err != nil {
		render.Render(w, r, RenderInvalidRequestError(err))
		return
	}

	// Update the item in database with the decoded updatedItemData given its itemId
	item, err := databaseInstance.UpdateItem(itemId, updatedItemData)

	if err != nil {
		if err == db.ErrorNoMatch {
			render.Render(w, r, ErrorNotFound)
		} else {
			render.Render(w, r, RenderServerError(err))
		}
		return
	}

	// Return the updated item to the client to indicate that the request was successful
	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, RenderServerError(err))
	}
}

// Delete a bucket list item given its id
func deleteItem(w http.ResponseWriter, r *http.Request) {
	// Get the item id URL parameter from the request context
	itemId := r.Context().Value(itemIDKey).(int)

	// delete the item in database with the provided id
	err := databaseInstance.DeleteItem(itemId)

	if err != nil {
		if err == db.ErrorNoMatch {
			render.Render(w, r, ErrorNotFound)
		} else {
			render.Render(w, r, RenderServerError(err))
		}
	}
}
