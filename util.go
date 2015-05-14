package main

import (
    "net/http"

    // "strings"
    // "encoding/json"
    // "fmt"
)

type Response struct {
    Status      string      `json:"status"`
    Data        interface{} `json:"data"`
}

type ResponseErr struct {
    Status      string      `json:"status"`
    Emsg        string      `json:"emsg"`
}

func finishOk(w http.ResponseWriter, data interface{}) {
    rjson.JSON(w, http.StatusOK, Response{Status:"ok", Data:data})
}

func finishErr(w http.ResponseWriter, emsg string) {
    rjson.JSON(w, http.StatusNotFound, ResponseErr{Status:"err", Emsg: emsg})
}
