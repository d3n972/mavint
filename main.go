package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/d3n972/mavint/models"

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
		Root:      "templates",
		Extension: ".tmpl",
		Master:    "layouts/master",
		Funcs: template.FuncMap{
			"timediff": func(station models.TD_Scheduler) string {
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
				if a < 0 {
					return true
				}
				if a >= b && a < c {
					return true
				}
				return false
			},
			"timediffMins": func(station models.TD_Scheduler) float64 {
				if station.Arrive != nil && station.ActualOrEstimatedArrive != nil {
					t1 := *station.Arrive
					t2 := station.ActualOrEstimatedArrive
					delta := t2.Sub(t1)
					fmt.Sprintln("delta: %d", delta)
					return delta.Minutes()
				}
				return 0
			},
			"delayReasons": func(s models.TS_TrainSchedDetails) string {
				return strings.Join(s.Train.HavarianInfok.HavariaInfo, " ")
			},
			"getExpectedHHMM": func(t time.Time) string {
				return t.Format("15 óra 04 perc")
			},
			"isTrainDeparted": func(t time.Time) bool {
				now := time.Now()
				if now.Hour() > 12 && t.Hour() < 6 {
					t = t.AddDate(0, 0, 1)
				}
				fmt.Println(t.Format(time.RFC3339))
				return now.After(t)
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
			"loctime": func(x *time.Time) string {
				if x != nil {
					t := x.Local()
					return t.Format("15:04")
				}
				return ""
			},
			"getTrainName": func(x any) string {
				if x.(models.Scheduler).GetName() == nil {
					fmt.Printf("Code: %s, FType: %s\n", x.(models.Scheduler).GetCode(), x.(models.Scheduler).GetFullShortType())
					return x.(models.Scheduler).GetCode() + " " + x.(models.Scheduler).GetFullShortType()
				}
				return x.(models.Scheduler).GetCode() + " " + *x.(models.Scheduler).GetName() + " " + x.(models.Scheduler).GetFullShortType()
			},
		},
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
