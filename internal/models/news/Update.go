package news

import (
	"elastic-search/pkg/database/mysql"
	"errors"
)

func (n *News) Update(data News) error {
	db := mysql.Connect()
	res := db.Model(&n).Updates(data)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}
