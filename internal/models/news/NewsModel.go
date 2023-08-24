package news

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type News struct {
    gorm.Model
    Title string `json:"title"`
    Content string `json:"content"`
}

func (n News) GetId() string {
    return fmt.Sprintf("%d", n.ID)
}

func (n News) ToJson() string {
    j, _ := json.Marshal(n)

    return string(j)
}