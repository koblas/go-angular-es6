package main

import (
    "github.com/gin-gonic/gin"
    // "net/http"

    "strings"
    // "strconv"
    // "encoding/json"
    // "fmt"
)

//  When sent via JSON as a POST/PUT
type BodyEntry struct {
    Title  *string      `json:"title,omitempty"`
    Completed  *bool    `json:"completed,omitempty"`
}

var (
    todoEntries = NewTodoStore("db/todo.db")
)

//
//
//
func TodoGet(c *gin.Context) {
    idStr := c.Params.ByName("id")

    if len(idStr) == 0 {
        v := c.Request.Form.Get("completed")

        checkCompleted := false
        completed := false
        if v != "" {
            checkCompleted = true
            completed = strings.TrimSpace(v) == "true"
        }

        result := make([]TodoItem, 0)

        entries := todoEntries.entries.All()

        for  _, entry := range entries {
            ok := true
            if checkCompleted {
                ok = ok && (entry.Completed == completed)
            }

            if ok {
                result = append(result, *entry)
            }
        }

        finishOk(c, result)
    } else {
        entry := todoEntries.entries.Find(idStr)

        if entry != nil {
            finishOk(c, entry)
        } else {
            finishErr(c, "Not Found")
        }
    }
}

func TodoPost(c *gin.Context) {
    tentry := BodyEntry{}
    c.Bind(&tentry)

    title := ""
    if tentry.Title != nil {
        title = strings.TrimSpace(*tentry.Title)
    }

    if len(title) != 0 {
        entry := todoEntries.entries.Create(TodoItem{Title:title})

        todoEntries.Save()

        finishOk(c, entry)
    } else {
        finishErr(c, "Title is required")
    }
}

func TodoPut(c *gin.Context) {
    idStr := c.Params.ByName("id")

    entry := todoEntries.entries.Find(idStr)

    if entry != nil {
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

        todoEntries.Save()

        finishOk(c, *entry)
    } else {
        finishErr(c, "Not Found")
    }
}

func TodoDelete(c *gin.Context) {
    idStr := c.Params.ByName("id")

    if idStr != "" && todoEntries.entries.Delete(idStr) {
        todoEntries.Save()
        finishOk(c, nil)
        return
    }

    finishErr(c, "Not Found")
}
