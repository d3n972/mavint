package scheduledTasks

import (
	"context"
	"fmt"
	"time"
)

func GetRedisTask() Schedule {
	return Schedule{
		Interval: 1 * time.Second,
		Handler: func(ctx AppContext) {
			fmt.Printf("asdasd\n")
			t := time.Now()
			stat := ctx.Redis.Set(context.Background(), "foo", t.Unix(), 0)
			s, e := stat.Result()
			fmt.Printf("%s,%v", s, e)
		},
	}
}
