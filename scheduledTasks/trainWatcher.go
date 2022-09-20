package scheduledTasks

import (
	"context"
	"encoding/json"
	"fmt"

	"time"

	"github.com/d3n972/mavint/models"
	"github.com/d3n972/mavint/models/db"
)

func WatchTrainsTask() *Schedule {
	return &Schedule{
		Interval: 15 * time.Second,
		Handler: func(ctx AppContext) {
			_tz := time.Local
			time.LoadLocation("UTC")
			R := ctx.Redis
			if R.Exists(context.TODO(), "havariaCache").Val() != 0 {
				hc := models.HavariaCache{}
				cacheObject, _ := R.Get(context.TODO(), "havariaCache").Bytes()
				json.Unmarshal(cacheObject, &hc)
				trains := []db.WatchedTrain{}

				ctx.Db.Find(&trains, "watch_until >= ?", time.Now().UTC().Format(time.RFC3339))

				fmt.Printf("%+v", trains)
			}
			time.LoadLocation(_tz.String())
		},
	}
}
