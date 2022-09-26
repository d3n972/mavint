package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/d3n972/mavint/models"
)

func GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"hhmm": func(v float64) string {
			t := time.Time{}
			t = t.Add(time.Duration(v) * time.Minute)
			fmt.Printf("%+v\n%+v\n", t, time.Duration(v)*time.Minute)
			return t.Format("15 óra 04 perc")
		},
		"hhmmTime": func(v time.Time) string {
			return v.Format("15 óra 04 perc")
		},
		"json": func(v any) string {
			k, e := json.MarshalIndent(v, "", "    ")
			if e != nil {
				fmt.Errorf("%s", e.Error())
				return "{\"err\":" + e.Error() + "\"}"
			}
			return string(k)
		},
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
		"getServiceIcons": func(train models.Train) string {
			return train.GetIconCharacters()
		},
		"timediffMins": func(station models.Scheduler) time.Duration {
			if station.Arrive != nil && station.ActualOrEstimatedArrive != nil {
				t1 := *station.Arrive
				t2 := station.ActualOrEstimatedArrive
				delta := t2.Sub(t1)
				return delta
			}
			return 0 * time.Minute
		},
		"delayReasons": func(s models.TrainSchedulerDetails) string {
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

			if x.(models.Train).GetFullShortType() == "InterCity" {
				return "IC" + x.(models.Train).GetCode() + " " + *x.(models.Train).GetName()
			}
			if x.(models.Train).GetFullShortType() == "InterRégió" {
				return "IR" + x.(models.Train).GetCode() + " " + *x.(models.Train).GetName()
			}
			if x.(models.Train).GetFullShortType() == "railjet xpress" {
				return "RJX" + x.(models.Train).GetCode()
			}
			if x.(models.Train).GetFullShortType() == "EuroCity" {
				return "EC" + x.(models.Train).GetCode() + " " + *x.(models.Train).GetName()
			}
			if x.(models.Train).GetFullShortType() == "EuroNight" {
				return "EN" + x.(models.Train).GetCode() + " " + *x.(models.Train).GetName()
			}
			if x.(models.Train).GetFullShortType() == "szeméyvonat" {
				return x.(models.Train).GetCode() + " " + *x.(models.Train).GetName()
			}
			return x.(models.Train).GetCode() + " " + *x.(models.Train).GetName()
		},
		"fGetCSSByDelay": func(f float64) string {
			return CSSColByDelay(time.Duration(int64(f)) * time.Minute)
		},
		"getCSSByDelay": CSSColByDelay,
	}

}
func CSSColByDelay(d time.Duration) string {
	colorCode := ""
	if d > 0*time.Minute && d <= 2*time.Minute {
		colorCode = "#009f7b"
	} else if d > 2*time.Minute && d <= 5*time.Minute {
		colorCode = "#2dc73b"
	} else if d > 5*time.Minute && d <= 10*time.Minute {
		colorCode = "#b3de07"
	} else if d > 10*time.Minute && d <= 20*time.Minute {
		colorCode = "#eed202"
	} else if d > 20*time.Minute && d <= 30*time.Minute {
		colorCode = "#cea104"
	} else if d > 30*time.Minute && d <= 40*time.Minute {
		colorCode = "#c57f07"
	} else if d > 40*time.Minute && d <= 50*time.Minute {
		colorCode = "#c1570b"
	} else if d > 50*time.Minute && d <= 60*time.Minute {
		colorCode = "#b6100a"
	} else if d > 60*time.Minute {
		colorCode = "#6e0e0a"
	} else {
		colorCode = "#99ffdd"
	}
	return colorCode
}
