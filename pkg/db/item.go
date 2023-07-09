package db

//--
// File that holds interactions with the Items table in terms of CRUD ops with the item model.
//--

import (
	"database/sql"

	"github.com/Shayne3000/Buckets/pkg/models"
)

func (db Database) AddItem(item *models.Item) error {
	var id int

	var created_at string

	query := `INSERT INTO items (name, description) VALUES ($1 $2) RETURNING id, created_at;`

	err := db.connection.QueryRow(query, item.Name, item.Description).Scan(&id, &created_at)

	if err != nil {
		return err
	}

	item.ID = id
	item.CreatedAt = created_at

	return nil
}

func (db Database) GetAllItems() (*models.ItemList, error) {
	itemsList := &models.ItemList{}

	rows, err := db.connection.Query("SELECT * FROM items ORDER BY ID DESC;")

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

	row := db.connection.QueryRow(query, id)

	switch err := row.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt); err {
	case sql.ErrNoRows:
		return item, errorNoMatch

	default:
		return item, err
	}
}

func (db Database) UpdateItem(itemId int, itemData models.Item) (models.Item, error) {
	item := models.Item{}

	query := `UPDATE items SET name=$1, description=$2 WHERE id=$3 RETURNING id, name, description, created_at;`

	err := db.connection.QueryRow(query, itemData.Name, itemData.Description, itemId).Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return item, errorNoMatch
		}

		return item, err
	}

	return item, nil
}

func (db Database) DeleteItem(itemId int) error {
	query := `DELETE FROM items WHERE id=$1;`

	_, err := db.connection.Exec(query, itemId)

	switch err {
	case sql.ErrNoRows:
		return errorNoMatch
	default:
		return err
	}
}
