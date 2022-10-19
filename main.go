package main

import (
	"embed"
	"fmt"
	"github.com/artonge/go-gtfs"
	"github.com/d3n972/mavint/auth"
	"github.com/d3n972/mavint/controllers"
	"github.com/d3n972/mavint/db"
	M "github.com/d3n972/mavint/models/db"
	"github.com/d3n972/mavint/scheduledTasks"
	"github.com/d3n972/mavint/services"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v9"
	"net/http"
	"os"
	"path/filepath"
)

//go:embed assets/*
var Assets embed.FS

//go:embed templates/*
var Templates embed.FS

//go:embed public/*
var pwaManifest embed.FS

func globalRecover(c *gin.Context) {
	defer services.PanicHandler(c, nil)
	c.Next()
}
func embeddedFH(config goview.Config, tmpl string) (string, error) {
	path := filepath.Join(config.Root, tmpl)
	bytes, err := Templates.ReadFile(path + config.Extension)
	return string(bytes), err
}
func main() {
	appCtx := scheduledTasks.AppContext{}
	appCtx.Db = db.GetDbInstance()
	appCtx.Redis = redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", // no password set
		DB:       0,                                  // use default DB
	})

	sentry.Init(sentry.ClientOptions{
		Dsn: "https://c8d19ca3e2214dda92fd358c2d853029@gt.d3n.it/1",
	})

	// Since sentry emits events in the background we need to make sure
	// they are sent before we shut down

	//load gtfs
	if g, err := gtfs.Load("/var/lib/gtfs", nil); err != nil {
		y := scheduledTasks.GTFSUpdaterTask()
		y.Handler(appCtx)
		g, _ := gtfs.Load("/var/lib/gtfs", nil)
		appCtx.Gtfs = g
	} else {
		appCtx.Gtfs = g
	}
	schedRunner := scheduledTasks.NewTaskRunner()
	schedRunner.AddTask("havariaUpdaterTask", scheduledTasks.HavarianUpdaterTask())
	schedRunner.AddTask("trainWatchTask", scheduledTasks.WatchTrainsTask())
	schedRunner.AddTask("EngineLoggerTask", scheduledTasks.EngineLoggerTask())
	schedRunner.AddTask("GTFSUpdaterTask", scheduledTasks.GTFSUpdaterTask())
	schedRunner.AddTask("VPEUpdaterTask", scheduledTasks.VPELoggerTask())
	y := scheduledTasks.VPELoggerTask()
	go y.Handler(appCtx)
	go schedRunner.Start(appCtx)
	os.Setenv("TZ", "Europe/Budapest")

	r := gin.Default()
	r.TrustedPlatform = gin.PlatformCloudflare
	r.Use(gin.Logger())
	r.Use(globalRecover)

	r.Use(func(ctx *gin.Context) {
		ctx.Set("cache", appCtx.Redis)
		ctx.Set("appctx", appCtx)

		ctx.Next()
	})
	r.Use(auth.SessionMiddleware)
	gvEngine := ginview.New(goview.Config{
		Root:         "templates",
		Extension:    ".tmpl",
		Master:       "layouts/master",
		Funcs:        GetFuncMap(),
		DisableCache: true,
		Delims:       goview.Delims{Left: "{{", Right: "}}"},
	})
	if os.Getenv("GIN_MODE") == "release" {
		fmt.Println("[Running in release mode, using embedded templates]")
		gvEngine.ViewEngine.SetFileHandler(embeddedFH)
	}
	r.HTMLRender = gvEngine

	r.StaticFS("/public", http.FS(Assets))
	r.StaticFS("/app", http.FS(pwaManifest))
	ttblCtrl := controllers.TimetableController{}
	tdCtrl := controllers.TrainDetailsController{}
	ticketCtrl := controllers.TicketController{}
	mapController := controllers.MapController{}
	emigController := controllers.EmigController{}
	newsController := controllers.NewsController{}
	twController := controllers.TrainWatchController{}
	edController := controllers.EngineDetailsController{}

	r.GET("/ed", edController.CountsForDay)
	r.GET("/ed/:date/:uic", edController.Render)

	r.GET("/emig", emigController.Render)
	r.GET("/auth/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "auth/login", gin.H{})
	})
	r.POST("/auth/login", func(c *gin.Context) {
		p, _ := c.GetPostForm("floatingPassword")
		u, _ := c.GetPostForm("floatingInput")
		c.JSON(http.StatusOK, gin.H{
			"p": p,
			"u": u,
		})
	})
	r.GET("/getdata/emig", emigController.GetTrainEngines)
	r.GET("/tt", ttblCtrl.Render)
	r.GET("/watch/form", twController.Render)
	r.POST("/watch/add", twController.Save)
	r.GET("/station/:station_code", ttblCtrl.Render)
	r.GET("/station", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/station/00"+ctx.Query("station_id"))
	})
	r.GET("/station_select", ttblCtrl.RenderSelectorPage)
	r.GET("/m", tdCtrl.Render)
	r.GET("/esd", tdCtrl.ESDDisplay)
	r.GET("/news", newsController.Render)
	r.GET("/article", newsController.RenderArticle)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pages/index", gin.H{})
	})
	r.GET("/map", mapController.Render)
	r.GET("/map/getdata", mapController.GetData)

	r.GET("/train/:train", tdCtrl.Render)
	r.GET("/ticket", ticketCtrl.Render)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	appCtx.Db.AutoMigrate(M.WatchedTrain{})

	r.Run(":12700") // listen and serve on 0.0.0.0:12700
}
