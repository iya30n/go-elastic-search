package news

import (
	"elastic-search/pkg/database/mysql"
	"errors"
)

func (n *News) Find(id string) error {
    db := mysql.Connect()
    db.Find(&n, "id", id)

    if n.ID == 0 {
        return errors.New("news not found")
    }

    return nil
}
