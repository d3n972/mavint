package main

import (
	"embed"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/d3n972/mavint/controllers"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var templates embed.FS

//go:embed assets/*
var assets embed.FS

func main() {
	r := gin.Default()

	r.HTMLRender = ginview.New(goview.Config{
		Root:      "templates",
		Extension: ".tmpl",
		Master:    "layouts/master",
		Funcs: template.FuncMap{
			"getColorOrFallback": func(a *string, b *string) string {
				if a == nil {
					return *b
				}
				return *a
			},
			"toIconName": func(trainType string) string {
				t := strings.ToLower(trainType)
				switch t {
				case "euronight":
					return "euronight"
				case "expresszvonat":
					return "expresszvonat"
				case "gyorsvonat":
					return "gyorsvonat"
				case "intercity":
					return "intercity"
				case "interregio":
					return "interregio"
				case "railjet":
					return "railjet"
				case "sebesvonat":
					return "sebesvonat"
				case "személyvonat":
					return "szemelyvonat"
				case "euregio":
					return "euregio"
				case "eurocity":
					return "eurocity"
				}
				return "szemelyvonat"
			},
			"toTrainType": func(t string) string {
				switch t {
				case "BESZ":
					return ""
				case "BEG":
					return "G"
				case "BEZ":
					return "Z"
				default:
					return t
				}
			},
			"loctime": func(x time.Time) string {
				return x.Local().Format(time.Kitchen)
			},
		},
		DisableCache: true,
		Delims:       goview.Delims{Left: "{{", Right: "}}"},
	})

	r.StaticFS("/public", http.FS(assets))
	ttblCtrl := controllers.TimetableController{}
	tdCtrl := controllers.TrainDetailsController{}
	r.GET("/tt", ttblCtrl.Render)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	r.GET("/train/:train_id", tdCtrl.Render)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":12700") // listen and serve on 0.0.0.0:12700
}
