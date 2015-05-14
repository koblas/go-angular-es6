package main

import (
  "github.com/julienschmidt/httprouter"
  "net/http"
  // "fmt"
)

func appIndex(w http.ResponseWriter, req *http.Request) {
    http.ServeFile(w, req, "./static/app.html")
}

func InitRouter() *httprouter.Router {
    router := httprouter.New()

    // Everything brings up the app
    router.NotFound = appIndex

    router.POST("/api/v1/auth/", AuthPost)

    router.GET("/api/v1/todo/", TodoGet)
    router.GET("/api/v1/todo/:id", TodoGet)
    router.POST("/api/v1/todo/", TodoPost)
    router.PUT("/api/v1/todo/:id", TodoPut)
    router.DELETE("/api/v1/todo/:id", TodoDelete)

    return router
}
