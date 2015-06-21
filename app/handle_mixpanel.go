package app

import (
    "github.com/gin-gonic/gin"
    "encoding/base64"
    "encoding/json"
    "net/http"

    // "strings"
    "fmt"
    // "bytes"
    // "os"
    // "log"
)

///
///
var (
    backendChannel = make(chan mixpanelEvent, 1000)
)

func consumeMessage() {
    for item := range backendChannel {
        fmt.Println("Got item", item)
    }
}

///
///

func initMixpanel() {
    go consumeMessage()
}

type MixpanelService struct {
    app     *Application
}

type mixpanelEvent struct {
    Event       string      `json:"event"`
    Properties  map[string]interface{}      `json:"properties"`
}

func sendResponse(c *gin.Context, status int, emsg string, verbose bool) {
    if verbose {
        c.JSON(http.StatusOK, map[string]interface{}{"status":status, "error":emsg})
    } else {
        c.JSON(http.StatusOK, status)
    }
}

func (svc *MixpanelService) Track(c *gin.Context) {
    form_data := c.Request.FormValue("data")

    verbose := c.Request.FormValue("verbose") == "1"

    if form_data == "" {
        sendResponse(c, 0, "Missing required argument data", verbose)
        return
    }

    bdata, err := base64.StdEncoding.DecodeString(form_data)

    if err != nil {
        sendResponse(c, 0, "Bad base64 data payload", verbose)
        return
    }

    event := mixpanelEvent{}

    err = json.Unmarshal(bdata, &event)
    if err != nil {
        sendResponse(c, 0, "Bad json data payload", verbose)
        return
    }

    fmt.Println(event)

    //
    // TODO -
    //   Check Token
    //   Check that event is not ""
    //   add "time" if missing
    //   

    /*
    b, err := json.Marshal(&event)

    var out bytes.Buffer
    json.Indent(&out, b, "", "  ")
    out.WriteTo(os.Stdout)
    */

    backendChannel <- event

    sendResponse(c, 1, "", verbose)
}
