package main

import (
  "github.com/codegangsta/negroni"
  "github.com/unrolled/render"
  "net/http"
  "fmt"
  "flag"
)

var (
    inDebug = flag.Bool("debug", false, "Run in debug mode")
    port    = flag.Int("port", 3000, "Port to serve on")
)

var (
    rjson = render.New(render.Options{
            IndentJSON: true,
    })
)

func main() {
    router := InitRouter()

    n := negroni.New(negroni.NewRecovery(), negroni.NewLogger(), &negroni.Static{Dir: http.Dir("static"), Prefix: "/static", IndexFile: ""})
    n.UseHandler(router)
    n.Run(fmt.Sprintf(":%d", *port))
}
