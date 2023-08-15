package main

import (
	"elastic-search/internal/models/news"
	"elastic-search/pkg/database/mysql"
	"fmt"
)

func main() {
    migrations := map[string]interface{}{
        "news": news.NewsModel{},
    }

    db := mysql.Connect()

    for name, migration := range migrations {
        fmt.Printf("migrating %s \n", name)
        db.AutoMigrate(migration)
    }

    fmt.Println("Done!")
}
