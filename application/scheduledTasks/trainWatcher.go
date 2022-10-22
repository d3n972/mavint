package scheduledTasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/d3n972/mavint/domain"
	"github.com/d3n972/mavint/domain/models"
	services2 "github.com/d3n972/mavint/infrastructure/services"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type TrainWatchTask struct {
	interval time.Duration
}

func (t TrainWatchTask) Handler(ctx domain.AppContext) {
	_tz := time.Local
	time.LoadLocation("UTC")
	R := ctx.Redis
	if R.Exists(context.TODO(), "havariaCache").Val() != 0 {
		fmt.Println("we have havaria\n")
		hc := models.HavariaCache{}
		cacheObject, _ := R.Get(context.TODO(), "havariaCache").Bytes()
		json.Unmarshal(cacheObject, &hc)
		trains := []db.WatchedTrain{}
		ctx.Db.Find(&trains, "watch_until >= ?", time.Now().UTC().Format(time.RFC3339))
		fmt.Printf("trains: %+v\n", trains)
		for _, train := range trains {
			fmt.Printf("R.Exists(context.TODO(), \"trainwatch:delay:\"+train.TrainID).Val(): %v", R.Exists(context.TODO(), "trainwatch:delay:"+train.TrainID).Val())
			if R.Exists(context.TODO(), "trainwatch:delay:"+train.TrainID).Val() > 0 {
				oldDelay, _ := R.Get(context.TODO(), "trainwatch:delay:"+train.TrainID).Int()
				fmt.Printf("oldDelay < hc[train.TrainID].Time: %v", oldDelay < hc[train.TrainID].Time)
				if oldDelay < hc[train.TrainID].Time {
					t.sendNotice(hc, train)
					R.Set(context.TODO(), "trainwatch:delay:"+train.TrainID, hc[train.TrainID].Time, 0)
				}
			} else {
				R.Set(context.TODO(), "trainwatch:delay:"+train.TrainID, hc[train.TrainID].Time, 0)
				t.sendNotice(hc, train)
			}
		}
	}
	time.LoadLocation(_tz.String())
}
func (t TrainWatchTask) sendNotice(hc models.HavariaCache, train db.WatchedTrain) {
	notificationService := services2.DiscordNotification{}
	notificationService.Init(
		&services2.Params{
			"id":    "1021878794790907944",
			"token": "YoBNiLmI2PS8pDlIJatOx-osh25iDu5S237awOXpF81qjoFehKUS5v6Hhl6x-x0Gq_Md",
			"name":  "TrainNotice",
		})
	templateS := `**VI_Train Delay Notice**
The VI_Train with number **{{.TrainID}}** is running **{{.Delay}}** minutes late!
`
	tmpl, _ := template.New("delaynotice").Parse(templateS)
	b := new(strings.Builder)
	tmpl.Execute(b, map[string]string{
		"TrainID": train.TrainID,
		"Delay":   strconv.Itoa(hc[train.TrainID].Time),
	})
	e := notificationService.Send(b.String())
	fmt.Printf("[e] %v\n\n", e)
}
