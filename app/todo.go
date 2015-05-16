package app

import (
	"github.com/nu7hatch/gouuid"
    "time"
)

func (item *TodoItem) setId() {
	id, _ := uuid.NewV4()
	item.Id = id.String()
}

type TodoItem struct {
	Id        string `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Order     int    `json:"order"`
	Text      string `json:"text"`
    CreatedAt   time.Time
}
