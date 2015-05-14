package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func HomeHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    rjson.JSON(w, http.StatusOK, map[string]string{"welcome": "This is rendered JSON!"})
}

func PageHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
    val := ps.ByName("name")

    rjson.JSON(w, http.StatusOK, map[string]string{"welcome": "This is rendered JSON!", "id": val})
}
