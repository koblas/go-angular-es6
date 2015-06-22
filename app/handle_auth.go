package app

import (
    "github.com/gin-gonic/gin"

    "fmt"
    "strings"
)

type UserEntry struct {
    Id          int         `json:"id"`
    Email       string      `json:"email"`
    Username    string      `json:"username"`
    Password    string      `json:"password"`
}

//  When sent via JSON as a POST/PUT
type bodyEntry struct {
    Email       *string      `json:"email,omitempty"`
    Password    *string      `json:"password,omitempty"`
    Token       *string      `json:"token,omitempty"`
    Params struct {
        Username    *string      `json:"username,omitempty"`
    } `json:"params"`
}

type AuthService struct {
    app     *Application
}

var (
    users = make([]UserEntry, 0)
    userIndex   = 0
)

//
func getUserById(id int) *UserEntry {
    for idx, e := range users {
        if e.Id == id {
            return &users[idx]
        }
    }

    return nil
}

//
//
//
func login() string {
    return "test"
}

func (svc *AuthService) registerHandler(c *gin.Context) {
    data := bodyEntry{}
    c.Bind(&data)

    username := ""
    email    := ""
    password := ""

    if data.Email != nil {
        email = strings.TrimSpace(*data.Email)
    }
    if data.Params.Username != nil {
        username = strings.TrimSpace(*data.Params.Username)
    }
    if data.Password != nil {
        password = strings.TrimSpace(*data.Password)
    }

    if len(email) == 0 || len(username) == 0 || len(password) == 0 {
        finishErr(c, "Missing argument - username, email or password")
        return
    }

    tuser := User{}

    if ! svc.app.db.Where(User{Email: email}).First(&tuser).RecordNotFound() {
        finishErr(c, "Existing User")
        return
    }

    user := NewUser(email, username)
    user.setPassword(password)

    err := svc.app.db.Create(&user).Error
    if err != nil {
        fmt.Println(err)
        finishErr(c, "Existing User")
        return
    }

    token := user.getToken()

    finishOk(c, map[string]string{"token":token})
}

func (svc *AuthService) loginHandler(c *gin.Context) {
    data := bodyEntry{}
    c.Bind(&data)

    var token string

    email    := ""
    password := ""

    if data.Email != nil {
        email = strings.TrimSpace(*data.Email)
    }
    if data.Password != nil {
        password = strings.TrimSpace(*data.Password)
    }

    if data.Token != nil {
        // TODO: Lookup user by token
    } else if len(email) == 0 || len(password) == 0 {
        finishErr(c, "Email/Password doesn't match")
        return
    } else {
        user := User{}
        if svc.app.db.Where(User{Email: email}).First(&user).RecordNotFound() {
            finishErr(c, "Email/Password doesn't match")
            return
        }

        if ! user.validate(password) {
            finishErr(c, "Email/Password doesn't match")
            return
        }

        token = user.getToken()
    }

    if len(token) == 0 {
        finishErr(c, "Unkown error")
        return
    }

    finishOk(c, map[string]string{"token":token})
}

func (svc *AuthService) AuthPost(c *gin.Context) {
    if c.Request.FormValue("register") != "" {
        svc.registerHandler(c)
    } else {
        svc.loginHandler(c)
    }
}
