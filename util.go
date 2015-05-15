package main

import (
    "github.com/gin-gonic/gin"
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

func finishOk(c *gin.Context, data interface{}) {
    // rjson.JSON(w, http.StatusOK, Response{Status:"ok", Data:data})
    c.JSON(http.StatusOK, gin.H{"status":"ok", "data":data})
}

func finishErr(c *gin.Context, emsg string) {
    // rjson.JSON(w, http.StatusNotFound, ResponseErr{Status:"err", Emsg: emsg})
    c.JSON(http.StatusNotFound, gin.H{"status":"err", "emsg":emsg})
}
