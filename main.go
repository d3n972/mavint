package main

import (
	"embed"
	"net/http"

	"github.com/d3n972/mavint/controllers"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

//go:embed assets/*
var assets embed.FS

//go:embed public/*
var pwaManifest embed.FS

func main() {
	r := gin.Default()
	r.TrustedPlatform = gin.PlatformCloudflare

	r.HTMLRender = ginview.New(goview.Config{
		Root:         "templates",
		Extension:    ".tmpl",
		Master:       "layouts/master",
		Funcs:        GetFuncMap(),
		DisableCache: true,
		Delims:       goview.Delims{Left: "{{", Right: "}}"},
	})

	r.StaticFS("/public", http.FS(assets))
	r.StaticFS("/app", http.FS(pwaManifest))
	ttblCtrl := controllers.TimetableController{}
	tdCtrl := controllers.TrainDetailsController{}
	ticketCtrl := controllers.TicketController{}
	mapController := controllers.MapController{}
	r.GET("/tt", ttblCtrl.Render)
	r.GET("/station/:station_code", ttblCtrl.Render)
	r.GET("/m", tdCtrl.Render)
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
