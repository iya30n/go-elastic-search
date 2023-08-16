package news

import "gorm.io/gorm"

type News struct {
    gorm.Model
    Title string `json:"title"`
    Content string `json:"content"`
}

