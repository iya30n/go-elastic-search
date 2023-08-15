package news

import "gorm.io/gorm"

type NewsModel struct {
    gorm.Model
    title string
    content string
}

