package main

import (
	"embed"
	"fmt"
	"github.com/d3n972/mavint/application/auth"
	"github.com/d3n972/mavint/domain"
	"github.com/d3n972/mavint/infrastructure"
	"github.com/d3n972/mavint/infrastructure/controllers"
	"github.com/d3n972/mavint/infrastructure/db"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v9"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"
	"time"
)

//go:embed infrastructure/assets/*
var Assets embed.FS

//go:embed infrastructure/templates/*
var Templates embed.FS

//go:embed infrastructure/public/*
var pwaManifest embed.FS
var gt_dsn string

func globalRecover(c *gin.Context, err any) {
	if os.Getenv("GIN_MODE") != "release" {
		gt_dsn = "https://b622fdae7cdc49d3a036234ac3d0dfeb@gt.d3n.it/2"
	} else {
		gt_dsn = "https://c8d19ca3e2214dda92fd358c2d853029@gt.d3n.it/1"
	}
	sentry.Init(sentry.ClientOptions{
		Dsn:     gt_dsn,
		Release: Commit,
	})

	recovered := recover()
	if recovered != nil {
		sentry.CaptureException(recovered.(error))
	}
	if err != nil {
		sentry.CaptureException(err.(error))

	}
	if flusherr := sentry.Flush(2 * time.Second); !flusherr {
		panic("failed to log")
	}
	c.Next()
}
func embeddedFH(config goview.Config, tmpl string) (string, error) {
	path := filepath.Join(config.Root, tmpl)
	bytes, err := Templates.ReadFile(path + config.Extension)
	return string(bytes), err
}

var Commit = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}
	return "develop"
}()

func main() {
	appCtx := domain.AppContext{}
	appCtx.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Llongfile)
	appCtx.Db = db.GetDbInstance()
	appCtx.Redis = redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", // no password set
		DB:       0,                                  // use default DB
	})

	// Since sentry emits events in the background we need to make sure
	// they are sent before we shut down

	//load gtfs
	/*if g, err := gtfs.Load("/var/lib/gtfs", nil); err != nil {
		y := scheduledTasks.GTFSUpdaterTask()
		y.Handler(appCtx)
		g, _ := gtfs.Load("/var/lib/gtfs", nil)
		appCtx.Gtfs = g
	} else {
		appCtx.Gtfs = g
	}
	schedRunner := domain.NewTaskRunner()
	schedRunner.AddTask("havariaUpdaterTask", scheduledTasks.HavarianUpdaterTask())
	schedRunner.AddTask("trainWatchTask", scheduledTasks.TrainWatchTask{}.Handler)
	schedRunner.AddTask("EngineLoggerTask", scheduledTasks.EngineLoggerTask())
	schedRunner.AddTask("GTFSUpdaterTask", scheduledTasks.GTFSUpdaterTask())
	schedRunner.AddTask("VPEUpdaterTask", scheduledTasks.VPELoggerTask())
	y := scheduledTasks.VPELoggerTask()
	go y.Handler(appCtx)
	go schedRunner.Start(appCtx)
	os.Setenv("TZ", "Europe/Budapest")
	*/
	r := gin.New()
	r.TrustedPlatform = gin.PlatformCloudflare
	r.Use(gin.Logger())
	r.Use(gin.CustomRecovery(globalRecover))

	r.Use(func(ctx *gin.Context) {
		ctx.Set("cache", appCtx.Redis)
		ctx.Set("appctx", appCtx)

		ctx.Next()
	})
	r.Use(auth.SessionMiddleware)
	gvEngine := ginview.New(goview.Config{
		Root:         "infrastructure/templates",
		Extension:    ".tmpl",
		Master:       "layouts/master",
		Funcs:        infrastructure.GetFuncMap(),
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

	appCtx.Db.AutoMigrate(domain.WatchedTrain{})

	r.Run(":12700") // listen and serve on 0.0.0.0:12700
}
