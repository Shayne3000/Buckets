package models

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

// Bind will run after unmarshalling is complete. We
// do some post-processing/sanity checks here.
func (item *Item) Bind(r *http.Request) error {
	if item.Name == "" {
		return fmt.Errorf("name is a required field")
	}

	return nil
}

// Set the pointer type *Item to implement the Render interface from Chi.Render
func (item *Item) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Set the pointer type *ItemList to implement the Render interface from Chi.Render
func (itemList *ItemList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
