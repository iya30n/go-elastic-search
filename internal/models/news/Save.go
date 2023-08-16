package news

import "elastic-search/pkg/database/mysql"

func (n *News) Save() error {
    db := mysql.Connect()
    db.Create(&n)

    return nil
}
