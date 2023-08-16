package news

import "gorm.io/gorm"

type NewsModel struct {
    gorm.Model
    Title string
    Content string
}

