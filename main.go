package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/d3n972/mavint/controllers"
	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var templates embed.FS

//go:embed assets/*
var assets embed.FS

func main() {
	r := gin.Default()

	//templ := template.Must(template.New("").ParseFS(templates, "templates/*.tmpl"))

	templ := template.Must(template.New("").Funcs(template.FuncMap{
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
	}).ParseGlob("templates/*.tmpl"))
	fmt.Printf("templ.DefinedTemplates(): %+v\n", templ.DefinedTemplates())
	r.SetHTMLTemplate(templ)
	//r.LoadHTMLGlob("templates/*")
	// example: /public/assets/images/example.png
	r.StaticFS("/public", http.FS(assets))
	ttblCtrl := controllers.TimetableController{}
	tdCtrl := controllers.TrainDetailsController{}
	r.GET("/tt", ttblCtrl.Render)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	r.GET("/m", tdCtrl.Render)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":12700") // listen and serve on 0.0.0.0:12700
}
