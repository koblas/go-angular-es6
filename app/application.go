package app

import (
	"github.com/gin-gonic/gin"
	"github.com/koblas/likemark/conf"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	"fmt"
	"net/http"
)

type Application struct {
	Config conf.Config
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

    svc := TodoService{app: a}

	// Load the static assets
	router.Static("/static", "./static")

	router.POST("/api/v1/auth/", AuthPost)

	router.GET("/api/v1/todo", svc.TodoGet)
	router.GET("/api/v1/todo/:id", svc.TodoGet)
	router.POST("/api/v1/todo", svc.TodoPost)
	router.PUT("/api/v1/todo/:id", svc.TodoPut)
	router.DELETE("/api/v1/todo/:id", svc.TodoDelete)
}

func (a *Application) getDb() (gorm.DB, error) {
	// connectionString := cfg.DbUser + ":" + cfg.DbPassword + "@tcp(" + cfg.DbHost + ":3306)/" + cfg.DbName + "?charset=utf8&parseTime=True"
	// return gorm.Open("mysql", connectionString)

	db, err := gorm.Open("sqlite3", a.Config.Database)

    db.LogMode(true)
	db.SingularTable(true)

    return db, err
}

func (a *Application) Init() {
	a.engine = gin.Default()
	db, _ := a.getDb()

    a.db = db

	a.initRouter()
}

func (a *Application) Run() {
	a.engine.Run(fmt.Sprintf(":%d", a.Config.Port))
}

func (a *Application) Migrate() error {
	db, err := a.getDb()
	if err != nil {
		return err
	}

	db.AutoMigrate(&TodoItem{})

	return nil
}
