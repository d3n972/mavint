package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/d3n972/mavint/controllers"
	"github.com/d3n972/mavint/db"
	"github.com/d3n972/mavint/scheduledTasks"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v9"
)

//go:embed assets/*
var Assets embed.FS

//go:embed templates/*
var Templates embed.FS

//go:embed public/*
var pwaManifest embed.FS

func XHR(c *gin.Context) bool {
	return strings.ToLower(c.Request.Header.Get("X-Requested-With")) == "xmlhttprequest"
}
func globalRecover(c *gin.Context) {
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			// that recovery also handle XHR's
			// you need handle it
			if XHR(c) {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": rec,
				})
			} else {
				fmt.Printf("%+v\n", rec)
				fmt.Printf("%+v\n", reflect.TypeOf(rec))
				c.HTML(http.StatusOK, "500", gin.H{
					"error": rec.(error).Error(),
				})
			}
		}
	}(c)
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

	schedRunner := scheduledTasks.NewTaskRunner()
	schedRunner.AddTask("redisTask", scheduledTasks.GetRedisTask())
	go schedRunner.Start(appCtx)
	os.Setenv("TZ", "Europe/Budapest")

	r := gin.Default()
	r.TrustedPlatform = gin.PlatformCloudflare
	r.Use(gin.Logger())
	//r.Use(globalRecover)
	r.Use(func(ctx *gin.Context) {
		ctx.Set("cache", appCtx.Redis)
	})
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
	r.GET("/emig", emigController.Render)
	r.GET("/getdata/emig", emigController.GetTrainEngines)
	r.GET("/tt", ttblCtrl.Render)
	r.GET("/station/:station_code", ttblCtrl.Render)
	r.GET("/station", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/station/00"+ctx.Query("station_id"))
	})
	r.GET("/station_select", ttblCtrl.RenderSelectorPage)
	r.GET("/m", tdCtrl.Render)
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
	r.Run(":12700") // listen and serve on 0.0.0.0:12700
}
