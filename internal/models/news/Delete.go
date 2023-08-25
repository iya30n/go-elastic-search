package news

import (
	"elastic-search/pkg/database/mysql"
	"errors"
)

func (n *News) Delete() error {
	db := mysql.Connect()
	res := db.Delete(&n, n.ID)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}
