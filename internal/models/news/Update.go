package news

import "elastic-search/pkg/database/mysql"

func (n *News) Update(data News) error {
    db := mysql.Connect()
    db.Model(&n).Updates(data)

    return nil
}
