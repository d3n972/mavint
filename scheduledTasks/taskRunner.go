package scheduledTasks

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/go-redis/redis/v9"
)

type HandlerFunction func(ctx AppContext)

type AppContext struct {
	Db    interface{}
	Redis *redis.Client
}
type Schedule struct {
	Interval time.Duration
	LastRun  time.Time
	Handler  HandlerFunction
}

type TaskRunner struct {
	State  bool
	done   chan bool
	ticker *time.Ticker
	tasks  map[string]Schedule
}

func (t *TaskRunner) AddTask(name string, task Schedule) {
	t.tasks[name] = task
}
func (t *TaskRunner) RunTask(ctx AppContext) {
	defer func() {
		if err := recover(); err != nil {
			err := err.(error)
			fmt.Printf("[!] Recovered from panic: %s", err.Error())
			fmt.Printf("Trace: \n" + string(debug.Stack()))
		}
	}()
	select {
	case <-t.done:
		t.Stop()
	case tick := <-t.ticker.C:
		_ = tick
		fmt.Printf("tick: %d\n", tick.Unix())
		for name, s := range t.tasks {
			fmt.Println("Task:" + name)
			now := time.Now()
			tx := s.LastRun
			tx.Add(s.Interval)
			if tx.Before(now) {
				fmt.Printf("Running task: %s\n", name)
				s.Handler(ctx)

				s.LastRun = now
			}
		}
	}
}
func (t *TaskRunner) Start(appCtx AppContext) {
	go func(ctx AppContext) {
		for t.State {
			t.RunTask(ctx)
		}
	}(appCtx)
}
func (t *TaskRunner) Stop() {
	t.done <- true
	t.State = false
	t.ticker.Stop()
}

func NewTaskRunner() *TaskRunner {
	return &TaskRunner{
		done:   make(chan bool),
		ticker: time.NewTicker(10 * time.Second),
		tasks:  map[string]Schedule{},
		State:  true,
	}
}
