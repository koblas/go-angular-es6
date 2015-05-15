package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func finishOk(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, gin.H{"status":"ok", "data":data})
}

func finishErr(c *gin.Context, emsg string) {
    c.JSON(http.StatusNotFound, gin.H{"status":"err", "emsg":emsg})
}
