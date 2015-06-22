package app

import (
    "github.com/gin-gonic/gin"

    "strings"
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
    user := svc.app.GetCurrentUser(c.Request, nil)
    if user == nil {
        finishErr(c, "Not logged in")
        return
    }

    idStr := c.Params.ByName("id")

    if len(idStr) == 0 {
        v := c.Request.Form.Get("completed")

        var     entries []TodoItem

        query := svc.app.db.Where(&TodoItem{UserGuid: user.Guid})
        if v != "" {
            query = query.Where(&TodoItem{Completed: v == "true"})
        }

        query.Order("created_at asc").Find(&entries)

        if entries == nil {
            finishOk(c, []TodoItem{})
        } else {
            finishOk(c, entries)
        }
    } else {
        entry := TodoItem{Guid: idStr, UserGuid: user.Guid}

        if ! svc.app.db.First(&entry).RecordNotFound() {
            finishOk(c, entry)
        } else {
            finishErr(c, "Not Found")
        }
    }
}

func (svc *TodoService) TodoPost(c *gin.Context) {
    user := svc.app.GetCurrentUser(c.Request, nil)
    if user == nil {
        finishErr(c, "Not logged in")
        return
    }

    tentry := BodyEntry{}
    c.Bind(&tentry)

    title := ""
    if tentry.Title != nil {
        title = strings.TrimSpace(*tentry.Title)
    }

    if len(title) != 0 {
        entry := NewTodoItem()
        entry.Title = title
        entry.UserGuid = user.Guid

        svc.app.db.Create(&entry)

        finishOk(c, entry)
    } else {
        finishErr(c, "Title is required")
    }
}

func (svc *TodoService) TodoPut(c *gin.Context) {
    user := svc.app.GetCurrentUser(c.Request, nil)
    if user == nil {
        finishErr(c, "Not logged in")
        return
    }

    idStr := c.Params.ByName("id")

    entry := TodoItem{}

    if ! svc.app.db.Where(TodoItem{Guid: idStr, UserGuid: user.Guid}).First(&entry).RecordNotFound() {
        tentry := BodyEntry{}

        if c.Bind(&tentry) != nil {
            finishErr(c, "No Such Entry")
            return
        }

        if tentry.Title != nil {
            entry.Title = *tentry.Title
        }
        if tentry.Completed != nil {
            entry.Completed = *tentry.Completed
        }

        svc.app.db.Save(entry)

        finishOk(c, entry)
    } else {
        finishErr(c, "Not Found")
    }
}

func (svc *TodoService) TodoDelete(c *gin.Context) {
    user := svc.app.GetCurrentUser(c.Request, nil)
    if user == nil {
        finishErr(c, "Not logged in")
        return
    }

    idStr := c.Params.ByName("id")

    if idStr != "" {
        query := svc.app.db.Where(TodoItem{Guid: idStr, UserGuid: user.Guid}).Delete(TodoItem{})
        if query.RecordNotFound() {
            finishErr(c, "Not Found")
        } else {
            finishOk(c, nil)
        }
    } else {
        finishErr(c, "Not Found")
    }
}
