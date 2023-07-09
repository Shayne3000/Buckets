package models

//--
// Item model
//--

import (
	"fmt"
	"net/http"
)

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type ItemList struct {
	Items []Item `json:"items"`
}

// Set the pointer type *Item to implement the Binder interface from Chi.Render
// so as to be able to decode/unmarshall the request body into an Item struct before 
// any performing any DB ops on it.

// Bind will run after unmarshalling is complete. We
// can do some post-processing/sanity checks here.
func (item *Item) Bind(r *http.Request) error {
	if item.Name == "" {
		return fmt.Errorf("name is a required field")
	}

	return nil
}

// Set the pointer type *Item to implement the Renderer interface from Chi.Render
// so that we can use render to package/transform the Item struct as a JSON response to the client request.
func (item *Item) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Set the pointer type *ItemList to implement the Render interface from Chi.Render
func (itemList *ItemList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
