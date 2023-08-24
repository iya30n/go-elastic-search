package mysql

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var doOnce sync.Once
var singletonConnection *gorm.DB

func Connect() *gorm.DB {
	doOnce.Do(func() {
		dsn := "root:@tcp(127.0.0.1:3306)/go_elastic?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if err != nil {
			panic(err)
		}

        singletonConnection = db
	})

    return singletonConnection
}
