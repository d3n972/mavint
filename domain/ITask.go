package domain

import "time"

type ITask interface {
	Handler(ctx AppContext)
	GetInterval() time.Duration
}
