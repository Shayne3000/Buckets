package db

// file responsible for interacting with the Items table in terms of CRUD ops with the item model.
import (
	"github.com/Shayne3000/Buckets/pkg/models"
)

func (db Database) AddItem(item *models.Item) error {
	var id int

	var created_at string

	query := `INSERT INTO items (name, description) VALUES ($1 $2) RETURNING id, created_at`

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

	rows, err := db.connection.Query("SELECT * FROM items ORDER BY ID DESC")

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

func (db Database) GetItemById(id int) (*models.Item, error) {

}

func (db Database) UpdateItem(itemId int, itemData models.Item) (models.Item, error) {

}

func (db Database) DeleteItem(itemId int) error {

}
