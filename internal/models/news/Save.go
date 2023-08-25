package news

import "elastic-search/pkg/database/mysql"

func (n *News) Save() error {
	db := mysql.Connect()
	res := db.Create(&n)

	return res.Error
}
