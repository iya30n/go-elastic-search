package news

import (
	"elastic-search/pkg/database/mysql"
	"errors"
)

func (n *News) Find(id string) error {
	db := mysql.Connect()
	res := db.Find(&n, "id", id)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("news not found")
	}

	return nil
}
