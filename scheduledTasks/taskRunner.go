package scheduledTasks

import (
	"fmt"
	"time"
)

type HandlerFunction func()

type Schedule struct {
	Interval time.Duration
	LastRun  time.Time
	Handler  HandlerFunction
}

type TaskRunner struct {
	done   chan bool
	ticker *time.Ticker
	tasks  map[string]Schedule
}

func (t *TaskRunner) AddTask(name string, task Schedule) {
	t.tasks[name] = task
}
func (t *TaskRunner) RunTask() {
	select {
	case <-t.done:
		t.Stop()
	case tick := <-t.ticker.C:
		_ = tick
		for name, s := range t.tasks {
			now := time.Now()
			tx := s.LastRun
			tx.Add(s.Interval)
			if tx.Before(now) {
				fmt.Printf("Running task: %s\n", name)
				s.Handler()
				s.LastRun = now
			}
		}
	}
}
func (t *TaskRunner) Stop() {
	t.done <- true
	t.ticker.Stop()
}

func NewTaskRunner() *TaskRunner {
	return &TaskRunner{
		done:   make(chan bool),
		ticker: time.NewTicker(10 * time.Second),
		tasks:  map[string]Schedule{},
	}
}
