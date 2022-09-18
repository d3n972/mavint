package scheduledTasks

import (
	"context"
	"encoding/json"
	"github.com/d3n972/mavint/controllers"
	"github.com/d3n972/mavint/models"
	"strings"
	"time"
)

func HavarianUpdaterTask() *Schedule {
	return &Schedule{
		Interval: 1 * time.Minute,
		Handler: func(ctx AppContext) {
			m := controllers.MapController{}
			R := ctx.Redis
			if R.Exists(context.TODO(), "havariaCache").Val() != 222 {
				hc := models.HavariaCache{}
				trains := m.ApiGetTrainCoords().D.Result.Trains.Train
				for _, t := range trains {
					tn := strings.Replace(t.TrainNumber, "55", "", 1)
					hc[tn] = models.HavariaCacheEntry{Time: t.Delay}
				}
				b, e := json.Marshal(hc)
				if e != nil {
				}
				R.Set(context.TODO(), "havariaCache", b, 5*time.Minute)
			}
		},
	}
}
