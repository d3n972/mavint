package scheduledTasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/d3n972/mavint/services"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/d3n972/mavint/models"
	"github.com/d3n972/mavint/models/db"
)

func sendNotice(hc models.HavariaCache, train db.WatchedTrain) {
	notificationService := services.DiscordNotification{}
	notificationService.Init(
		// https://discord.com/api/webhooks/1021878794790907944/YoBNiLmI2PS8pDlIJatOx-osh25iDu5S237awOXpF81qjoFehKUS5v6Hhl6x-x0Gq_Md
		&services.Params{
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
func WatchTrainsTask() *Schedule {

	return &Schedule{
		Interval: 15 * time.Second,
		Handler: func(ctx AppContext) {
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
							sendNotice(hc, train)
							R.Set(context.TODO(), "trainwatch:delay:"+train.TrainID, hc[train.TrainID].Time, 0)
						}
					} else {
						R.Set(context.TODO(), "trainwatch:delay:"+train.TrainID, hc[train.TrainID].Time, 0)
						sendNotice(hc, train)
					}
				}
			}
			time.LoadLocation(_tz.String())
		},
	}
}
