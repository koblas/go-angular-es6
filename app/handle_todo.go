package app

import (
    "github.com/gin-gonic/gin"

    "strings"
    "fmt"
)

//  When sent via JSON as a POST/PUT
type BodyEntry struct {
    Title  *string      `json:"title,omitempty"`
    Completed  *bool    `json:"completed,omitempty"`
}

type TodoService struct {
    app     *Application
}

//
//
//
func (svc *TodoService) TodoGet(c *gin.Context) {
    idStr := c.Params.ByName("id")

    if len(idStr) == 0 {
        v := c.Request.Form.Get("completed")

        var     entries []TodoItem

        query := &svc.app.db
        if v != "" {
            query = query.Where(&TodoItem{Completed: v == "true"})
        }

        query.Order("created_at desc").Find(&entries)

        if entries == nil {
            finishOk(c, []TodoItem{})
        } else {
            finishOk(c, entries)
        }
    } else {
        entry := TodoItem{Id: idStr}

        if ! svc.app.db.First(&entry).RecordNotFound() {
            finishOk(c, entry)
        } else {
            finishErr(c, "Not Found")
        }
    }
}

func (svc *TodoService) TodoPost(c *gin.Context) {
    tentry := BodyEntry{}
    c.Bind(&tentry)

    title := ""
    if tentry.Title != nil {
        title = strings.TrimSpace(*tentry.Title)
    }

    if len(title) != 0 {
        entry := TodoItem{Title:title}
        entry.setId()
        id := entry.Id

        fmt.Println(entry)

        svc.app.db.Create(&entry)

        fmt.Println(entry)

        entry.Id = id

        finishOk(c, entry)
    } else {
        finishErr(c, "Title is required")
    }
}

func (svc *TodoService) TodoPut(c *gin.Context) {
    idStr := c.Params.ByName("id")

    entry := TodoItem{Id: idStr}

    if ! svc.app.db.First(&entry).RecordNotFound() {
        tentry := BodyEntry{}

        if !c.Bind(&tentry) {
            finishErr(c, "No Such Entry")
            return
        }

        if tentry.Title != nil {
            entry.Title = *tentry.Title
        }
        if tentry.Completed != nil {
            entry.Completed = *tentry.Completed
        }

        svc.app.db.Save(&entry)

        finishOk(c, entry)
    } else {
        finishErr(c, "Not Found")
    }
}

func (svc *TodoService) TodoDelete(c *gin.Context) {
    idStr := c.Params.ByName("id")

    if idStr != "" {
        query := svc.app.db.Delete(TodoItem{Id: idStr})
        if query.RecordNotFound() {
            finishErr(c, "Not Found")
        } else {
            finishOk(c, nil)
        }
    } else {
        finishErr(c, "Not Found")
    }
}
