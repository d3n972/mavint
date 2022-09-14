package main

import (
	"embed"
	"fmt"
	"github.com/d3n972/mavint/controllers"
	"github.com/d3n972/mavint/scheduledTasks"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v9"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

//go:embed assets/*
var Assets embed.FS

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

func main() {
	schedRunner := scheduledTasks.NewTaskRunner()
	schedRunner.AddTask("test", scheduledTasks.Schedule{
		Interval: 1 * time.Minute,
		Handler: func() {
			fmt.Printf("asdasd\n")
		},
	})
	go func() { schedRunner.RunTask() }()
	os.Setenv("TZ", "Europe/Budapest")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", // no password set
		DB:       0,                                  // use default DB
	})
	r := gin.Default()
	r.TrustedPlatform = gin.PlatformCloudflare
	r.Use(gin.Logger())
	//r.Use(globalRecover)
	r.Use(func(ctx *gin.Context) {

		ctx.Set("cache", rdb)
	})
	r.HTMLRender = ginview.New(goview.Config{
		Root:         "templates",
		Extension:    ".tmpl",
		Master:       "layouts/master",
		Funcs:        GetFuncMap(),
		DisableCache: true,
		Delims:       goview.Delims{Left: "{{", Right: "}}"},
	})

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
