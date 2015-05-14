package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"

    "strings"
    // "strconv"
    "encoding/json"
    // "fmt"
)

//  When sent via JSON as a POST/PUT
type BodyEntry struct {
    Title  *string      `json:"title,omitempty"`
    Completed  *bool    `json:"completed,omitempty"`
}

var (
    todoEntries = NewTodoStore("todo.db")
    todoOrder = 0
)

//
//
//
func TodoGet(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
    idStr := ps.ByName("id")

    if len(idStr) == 0 {
        v := req.FormValue("completed")

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

        finishOk(w, result)
    } else {
        entry := todoEntries.entries.Find(idStr)

        if entry != nil {
            finishOk(w, entry)
        } else {
            finishErr(w, "Not Found")
        }
    }
}

func TodoPost(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
    tentry := BodyEntry{}
    json.NewDecoder(req.Body).Decode(&tentry)

    title := ""
    if tentry.Title != nil {
        title = strings.TrimSpace(*tentry.Title)
    }

    if len(title) != 0 {
        entry := todoEntries.entries.Create(TodoItem{Title:title, Order: todoOrder})
        todoOrder++

        todoEntries.Save()

        finishOk(w, entry)
    } else {
        finishErr(w, "Title is required")
    }
}

func TodoPut(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
    idStr := ps.ByName("id")

    entry := todoEntries.entries.Find(idStr)

    if entry != nil {
        tentry := BodyEntry{}

        err := json.NewDecoder(req.Body).Decode(&tentry)

        if err != nil {
            rjson.JSON(w, http.StatusNotFound, ResponseErr{Status:"err", Emsg: "No Such Entry"})
            return
        }

        if tentry.Title != nil {
            entry.Title = *tentry.Title
        }
        if tentry.Completed != nil {
            entry.Completed = *tentry.Completed
        }

        todoEntries.Save()

        finishOk(w, *entry)
    } else {
        finishErr(w, "Not Found")
    }
}

func TodoDelete(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
    id := ps.ByName("id")

    if id != "" && todoEntries.entries.Delete(id) {
        todoEntries.Save()
        finishOk(w, nil)
        return
    }

    finishErr(w, "Not Found")
}
