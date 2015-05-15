package main

import (
	"github.com/nu7hatch/gouuid"
    "sort"
    "gopkg.in/yaml.v2"
    "fmt"
    "io/ioutil"
)

func newId() string {
	id, _ := uuid.NewV4()
	return id.String()
}

type TodoItem struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Order     int    `json:"order"`
	Text      string `json:"text"`
}

func (i *TodoItem) Update(item TodoItem) *TodoItem {
	i.Title = item.Title
	i.Completed = item.Completed
	i.Order = item.Order
	i.Text = item.Text
	return i
}

type Todo map[string]*  TodoItem

type TodoStore struct {
    file    string
    entries Todo
}

type TodoList []*TodoItem

//
func (s TodoList) Len() int {
    return len(s)
}

func (s TodoList) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s TodoList) Less(i, j int) bool {
    return s[i].Order < s[j].Order
}

///
//
func NewTodoStore(file string) TodoStore {
    d, err := ioutil.ReadFile(file)

    if err != nil {
        fmt.Println("Fatal", err)
        return TodoStore{}
    }

    entries := Todo{}
    err = yaml.Unmarshal(d, &entries)

    if err != nil {
        fmt.Println("Fatal Unmarshal", err)
        return TodoStore{}
    }

    return TodoStore{file: file, entries: entries}
}

func (t TodoStore) Save() {
    d, err := yaml.Marshal(t.entries)

    if err != nil {
        fmt.Println("Fatal", err)
        return
    }

    if t.file != "" {
        err = ioutil.WriteFile(t.file, d, 0644)
    }

    if err != nil {
        fmt.Println("Fatal Write", err)
        return
    }
}

///

func (t Todo) All() TodoList {
	items := TodoList{}
	for _, item := range t {
		items = append(items, item)
	}
    sort.Sort(items)
	return items
}

func (t Todo) Find(id string) *TodoItem {
	for _, item := range t {
		if item.Id == id {
			return item
		}
	}

	return nil
}

func (t Todo) Create(item TodoItem) *TodoItem {
    order := 0
	for _, item := range t {
        if item.Order > order {
            order = item.Order
        }
    }
	item.Id = newId()
    item.Order = order + 1
	t[item.Id] = &item
	return &item
}

func (t Todo) Update(id string, updatedItem TodoItem) *TodoItem {
	if item := t.Find(id); item != nil {
		return item.Update(updatedItem)
	} else {
		return nil
	}
}

func (t Todo) DeleteAll() string {
	for k := range t {
		delete(t, k)
	}
	return ""
}

func (t Todo) Delete(id string) bool {
	for k := range t {
		if k == id {
			delete(t, k)
            return true
		}
	}
	return false
}
