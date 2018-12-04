package main

import (
	"flag"
	"math/rand"
	"net/http"
	"os"
	"restdoc-api/app"
	"restdoc-api/tasks"
	"restdoc-api/web"
	"restdoc-api/web/middleware"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/panjiang/golog"
	"github.com/panjiang/overseer"

	_ "image/jpeg"
	_ "image/png"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var fork = flag.Bool("fork", false, "")

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// 初始化
func init() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	must(app.ParseConfig())
}

func initCom() {
	must(app.InitDB(app.Config.Mysql))
	must(app.InitRedisCli(app.Config.Redis))
}

// 运行任务
func runTasks() {
	tasks.DatabaseMigrate()
}

// 运行网站程序
func runWeb(state overseer.State) {
	r := gin.New()
	if app.Config.Release {
		gin.SetMode(gin.ReleaseMode)
	} else {
		r.Use(gin.Logger())
	}

	r.Use(middleware.Recovery())
	// r.Use(middleware.CORS())
	r.Use(middleware.Session())

	web.Route(r)

	s := &http.Server{
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Info("Listening", app.Config.Bind, state.ID)
	s.Serve(state.Listener)
}

func prog(state overseer.State) {
	initCom()
	runTasks()
	runWeb(state)
}

func main() {
	if *fork {
		log.Info("Fork PID:", os.Getpid())
	} else {
		log.Info("Master PID:", os.Getpid())
	}

	err := overseer.RunErr(overseer.Config{
		Debug:   false,
		Program: prog,
		Address: app.Config.Bind,
		ForkArg: "-fork",
	})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Shutdown")
	}
}
