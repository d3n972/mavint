package main

import (
	"html/template"
	"strings"
	"time"

	"github.com/d3n972/mavint/models"
)

func GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"timediff": func(station models.TD_Scheduler) string {
			if station.Arrive != nil && station.ActualOrEstimatedArrive != nil {
				t1 := *station.Arrive
				t2 := *station.ActualOrEstimatedArrive
				delta := t2.Sub(t1)
				if delta.Minutes() > 0 {
					strDelay := time.Time{}.Add(delta).Format("15 óra 04 perc")
					return "+" + strings.Replace(strDelay, "00 óra ", "", -1)
				}
			}
			return ""
		},
		"toPrintPage": func(url string) string {
			parts := strings.Split(url, "/")
			return parts[len(parts)-1]
		},
		"iDelayInRange": func(a int, b int, c int) bool {
			if a >= b && a < c {
				return true
			}
			return false
		},
		"delayInRange": func(a float64, b float64, c float64) bool {
			//fmt.Printf("float64[a, b, c]: %v\n", []float64{a, b, c})

			if a >= b && a < c {
				return true
			}
			return false
		},
		"getServiceIcons": func(train models.Scheduler) string {
			return train.GetIconCharacters()
		},
		"timediffMins": func(station models.TD_Scheduler) float64 {
			if station.Arrive != nil && station.ActualOrEstimatedArrive != nil {
				t1 := *station.Arrive
				t2 := station.ActualOrEstimatedArrive
				delta := t2.Sub(t1)
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
			case "interrégió":
				return "interregio"
			case "railjet xpress":
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
				time.LoadLocation("Europe/Budapest")
				t := x.Local()
				return t.Format("15:04")
			}
			return ""
		},
		"getTrainName": func(x any) string {

			if x.(models.Scheduler).GetFullShortType() == "InterCity" {
				return "IC" + x.(models.Scheduler).GetCode() + " " + *x.(models.Scheduler).GetName()
			}
			if x.(models.Scheduler).GetFullShortType() == "InterRégió" {
				return "IR" + x.(models.Scheduler).GetCode() + " " + *x.(models.Scheduler).GetName()
			}
			if x.(models.Scheduler).GetFullShortType() == "railjet xpress" {
				return "RJX" + x.(models.Scheduler).GetCode()
			}
			if x.(models.Scheduler).GetFullShortType() == "EuroCity" {
				return "EC" + x.(models.Scheduler).GetCode() + " " + *x.(models.Scheduler).GetName()
			}
			if x.(models.Scheduler).GetFullShortType() == "EuroNight" {
				return "EN" + x.(models.Scheduler).GetCode() + " " + *x.(models.Scheduler).GetName()
			}
			if x.(models.Scheduler).GetFullShortType() == "szeméyvonat" {
				return x.(models.Scheduler).GetCode() + " " + *x.(models.Scheduler).GetName()
			}
			return x.(models.Scheduler).GetCode() + " " + *x.(models.Scheduler).GetName()
		},
	}

}
