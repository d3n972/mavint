package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/d3n972/mavint/controllers"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

//go:embed assets/*
var assets embed.FS

type QQ_1 struct {
	Arrive                  *time.Time `json:"arrive"`
	Start                   *time.Time `json:"start"`
	ActualOrEstimatedArrive *time.Time `json:"actualOrEstimatedArrive"`
	ActualOrEstimatedStart  *time.Time `json:"actualOrEstimatedStart"`
}

func main() {
	r := gin.Default()

	r.HTMLRender = ginview.New(goview.Config{
		Root:      "templates",
		Extension: ".tmpl",
		Master:    "layouts/master",
		Funcs: template.FuncMap{
			"timediff": func(station controllers.TD_Scheduler) string {
				if station.Arrive != nil && station.ActualOrEstimatedArrive != nil {
					t1 := *station.Arrive
					t2 := *station.ActualOrEstimatedArrive
					delta := t2.Sub(t1)
					if delta.Minutes() > 0 {
						strDelay := time.Time{}.Add(delta).Format("15 óra 04 perc")
						return "Késés: " + strings.Replace(strDelay, "00 óra ", "", -1)
					}
				}
				return ""
			},
			"delayInRange": func(a float64, b float64, c float64) bool {
				fmt.Printf("float64[a, b, c]: %v\n", []float64{a, b, c})
				if a > b && a < c {
					return true
				}
				return false
			},
			"timediffMins": func(station controllers.TD_Scheduler) float64 {
				if station.Arrive != nil && station.ActualOrEstimatedArrive != nil {
					t1 := *station.Arrive
					t2 := *station.ActualOrEstimatedArrive
					delta := t2.Sub(t1)
					return delta.Minutes()
				}
				return 0
			},
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
	r.GET("/m", tdCtrl.Render)
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
