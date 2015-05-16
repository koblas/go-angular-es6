package app

import (
    "github.com/gin-gonic/gin"

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
    Username    *string      `json:"username,omitempty"`
    Password    *string      `json:"password,omitempty"`
    Token       *string      `json:"token,omitempty"`
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

func registerHandler(c *gin.Context) {
    data := bodyEntry{}
    c.Bind(&data)

    username := ""
    email    := ""
    password := ""

    if data.Email != nil {
        email = strings.TrimSpace(*data.Email)
    }
    if data.Username != nil {
        username = strings.TrimSpace(*data.Username)
    }
    if data.Password != nil {
        password = strings.TrimSpace(*data.Password)
    }

    if len(email) == 0 || len(username) == 0 || len(password) == 0 {
        finishErr(c, "Missing argument - username, email or password")
        return
    }

    // check if unique email
    //    finishErr("Email already in use")

    // Create User

    // token := login(user)

    token := login()

    finishOk(c, map[string]string{"token":token})
}

func loginHandler(c *gin.Context) {
    data := bodyEntry{}
    c.Bind(&data)

    if data.Token != nil {
        // TODO: Lookup user by token

        // if user == nil 
    } else if data.Email == nil || data.Password == nil || len(*data.Email) == 0 || len(*data.Password) == 0 {
        finishErr(c, "Email/Password doesn't match")
        return
    } else {
    }

    // TODO: self.login(user)

    token := login()

    finishOk(c, map[string]string{"token":token})
}

func AuthPost(c *gin.Context) {
    if c.Request.Form.Get("register") != "" {
        registerHandler(c)
    } else {
        loginHandler(c)
    }
}
