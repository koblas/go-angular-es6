package main

import (
  "github.com/gin-gonic/gin"

  "net/http"
  // "fmt"
)

func appIndex(c *gin.Context) {
    http.ServeFile(c.Writer, c.Request, "./static/app.html")
}

func InitRouter(router *gin.Engine) {
    // Everything brings up the app
    router.NoRoute(appIndex)

    router.Static("/static", "./static")

    router.POST("/api/v1/auth/", AuthPost)

    router.GET("/api/v1/todo/", TodoGet)
    router.GET("/api/v1/todo/:id", TodoGet)
    router.POST("/api/v1/todo/", TodoPost)
    router.PUT("/api/v1/todo/:id", TodoPut)
    router.DELETE("/api/v1/todo/:id", TodoDelete)
}
