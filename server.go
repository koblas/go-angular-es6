package main

import (
  "github.com/gin-gonic/gin"
//  "net/http"
  "fmt"
  "flag"
)

var (
    inDebug = flag.Bool("debug", false, "Run in debug mode")
    port    = flag.Int("port", 3000, "Port to serve on")
)

func main() {
    engine := gin.Default()

    InitRouter(engine)

    engine.Run(fmt.Sprintf(":%d", *port))
}
