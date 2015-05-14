package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"

    "strings"
    "strconv"
    "encoding/json"
    // "fmt"
)

type TodoEntry struct {
    Id          int         `json:"id"`
    Title       string      `json:"title"`
    Completed   bool        `json:"completed"`
}

//  When sent via JSON as a POST/PUT
type BodyEntry struct {
    Title  *string      `json:"title,omitempty"`
    Completed  *bool    `json:"completed,omitempty"`
}

var (
    entries = make([]TodoEntry, 0)
    index   = 0
)

//
func getById(id int) *TodoEntry {
    for idx, e := range entries {
        if e.Id == id {
            return &entries[idx]
        }
    }

    return nil
}

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

        result := make([]TodoEntry, 0)

        for  _, entry := range entries {
            ok := true
            if checkCompleted {
                ok = ok && (entry.Completed == completed)
            }

            if ok {
                result = append(result, entry)
            }
        }

        finishOk(w, result)
    } else {
        id, err := strconv.Atoi(idStr)

        if err != nil {
            finishErr(w, "Bad ID")
            return
        }

        entry := getById(id)

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
        entry := TodoEntry{}
        index++
        entry.Id = index
        entry.Title = title
        entry.Completed = false

        entries = append(entries, entry)

        finishOk(w, entry)
    } else {
        finishErr(w, "Title is required")
    }
}

func TodoPut(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
    id, err := strconv.Atoi(ps.ByName("id"))

    if err != nil {
        finishErr(w, "Bad ID")
        return
    }

    entry := getById(id)

    if entry != nil {
        tentry := BodyEntry{}

        json.NewDecoder(req.Body).Decode(&tentry)

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

        finishOk(w, *entry)
    } else {
        finishErr(w, "Not Found")
    }
}

func TodoDelete(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
    id, err := strconv.Atoi(ps.ByName("id"))

    if err != nil {
        finishErr(w, "Bad ID")
        return
    }

    for idx, e := range entries {
        if e.Id == id {
            entries = append(entries[:idx], entries[idx+1:]...)
            finishOk(w, nil)
            return
        }
    }

    finishErr(w, "Not Found")
}
