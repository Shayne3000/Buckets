package db

//--
// File that holds CRUD interactions with the item resource in the Items table.
//--

import (
	"database/sql"

	"github.com/Shayne3000/Buckets/pkg/models"
)

func (db Database) AddItem(item *models.Item) error {
	var id int

	var created_at string

	query := `INSERT INTO items (name, description) VALUES ($1, $2) RETURNING id, created_at`

	err := db.Connection.QueryRow(query, item.Name, item.Description).Scan(&id, &created_at)

	if err != nil {
		return err
	}

	item.ID = id
	item.CreatedAt = created_at

	return nil
}

func (db Database) GetAllItems() (*models.ItemList, error) {
	itemsList := &models.ItemList{}

	// Todo Change the order to ascending
	rows, err := db.Connection.Query("SELECT * FROM items ORDER BY ID DESC;")

	if err != nil {
		return itemsList, err
	}

	for rows.Next() {
		var item models.Item

		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt)

		if err != nil {
			return itemsList, err
		}

		itemsList.Items = append(itemsList.Items, item)
	}

	return itemsList, nil
}

func (db Database) GetItemById(id int) (models.Item, error) {
	item := models.Item{}

	query := `SELECT * FROM items WHERE id = $1;`

	row := db.Connection.QueryRow(query, id)

	switch err := row.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt); err {
	case sql.ErrNoRows:
		return item, ErrorNoMatch

	default:
		return item, err
	}
}

func (db Database) UpdateItem(itemId int, itemData models.Item) (models.Item, error) {
	item := models.Item{}

	query := `UPDATE items SET name=$1, description=$2 WHERE id=$3 RETURNING id, name, description, created_at;`

	err := db.Connection.QueryRow(query, itemData.Name, itemData.Description, itemId).Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return item, ErrorNoMatch
		}

		return item, err
	}

	return item, nil
}

func (db Database) DeleteItem(itemId int) error {
	query := `DELETE FROM items WHERE id=$1;`

	_, err := db.Connection.Exec(query, itemId)

	switch err {
	case sql.ErrNoRows:
		return ErrorNoMatch
	default:
		return err
	}
}
