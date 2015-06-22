package app

import (
	"github.com/nu7hatch/gouuid"
    "time"
)

func generateId() string {
	id, _ := uuid.NewV4()
	return id.String()
}

type TodoItem struct {
	Id        int `json:"_" gorm:"primary_key"`
	Guid      string `json:"id" sql:"index"`
	UserGuid  string `json:"_" sql:"index"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Order     int    `json:"order"`
	Text      string `json:"text"`
    CreatedAt   time.Time
}

func NewTodoItem() *TodoItem {
    return &TodoItem{Guid: generateId()}
}
