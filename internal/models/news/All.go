package news

import (
	"elastic-search/pkg/database/mysql"
)

func (News) All() ([]*News, error) {
	news := []*News{}
	db := mysql.Connect()
	res := db.Find(&news)

	return news, res.Error
}
