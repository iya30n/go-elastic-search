package news

import "elastic-search/pkg/database/mysql"

func (News) All() ([]*News, error) {
    news := []*News{}
    db := mysql.Connect()
    db.Find(&news)

    return news, nil
}
