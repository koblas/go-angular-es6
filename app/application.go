package app

import (
	"github.com/gin-gonic/gin"
	"github.com/koblas/go-angular-es6/conf"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	"fmt"
	"net/url"
	"net/http"
)

type Application struct {
	Config *conf.ConfigData
	engine *gin.Engine
    db     gorm.DB
}

func appIndex(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "./static/app.html")
}

func (a *Application) initRouter() {
	router := a.engine

	// By Default - not found should be passed to Angular
	router.NoRoute(appIndex)

	// Load the static assets
	router.Static("/static", "./static")

    //
    //  Authentication handler
    //
    auth_svc := AuthService{app: a}

	router.POST("/api/v1/auth", auth_svc.AuthPost)

    // 
    //  Todo Handlers
    // 
    todo_svc := TodoService{app: a}

	router.GET("/api/v1/todo", todo_svc.TodoGet)
	router.GET("/api/v1/todo/:id", todo_svc.TodoGet)
	router.POST("/api/v1/todo", todo_svc.TodoPost)
	router.PUT("/api/v1/todo/:id", todo_svc.TodoPut)
	router.DELETE("/api/v1/todo/:id", todo_svc.TodoDelete)

    //
    //
    //
    mix_svc := MixpanelService{app: a}
	router.GET("/track", mix_svc.Track)
	router.GET("/track/", mix_svc.Track)
}

func (a *Application) getDb() (gorm.DB, error) {
	// connectionString := cfg.DbUser + ":" + cfg.DbPassword + "@tcp(" + cfg.DbHost + ":3306)/" + cfg.DbName + "?charset=utf8&parseTime=True"
	// return gorm.Open("mysql", connectionString)

	db, err := gorm.Open("sqlite3", a.Config.App.Database)

    db.LogMode(true)
	db.SingularTable(true)

    return db, err
}

func (a *Application) GetCurrentUser(req *http.Request, token *string) *User {

    var tokenStr string
    var err error

    if token == nil {
        cookie, err := req.Cookie(USER_COOKIE)
        if err != nil {
            return nil
        }
        tokenStr = cookie.Value
    } else {
        tokenStr = *token
    }

    if tokenStr[1] == '%' {
        tokenStr, err = url.QueryUnescape(tokenStr)
    }

    var guid string

    guid, err = DecodeSignedValue(SECRET, AUTH_NAME, tokenStr, nil)

    if err != nil {
        return nil
    }

    user := User{}
    if a.db.Where(User{Guid: guid}).First(&user).RecordNotFound() {
        return nil
    }

    return &user
}

func (a *Application) Init() {
	a.engine = gin.Default()
	db, _ := a.getDb()

    a.db = db

	a.initRouter()
    initMixpanel()      // TODO - this should be pushed into a mixpanel init
}

func (a *Application) Run() {
	a.engine.Run(fmt.Sprintf(":%d", a.Config.App.HttpPort))
}

func (a *Application) Migrate() error {
	db, err := a.getDb()
	if err != nil {
		return err
	}

	db.AutoMigrate(&TodoItem{})
	db.AutoMigrate(&User{})

	return nil
}
