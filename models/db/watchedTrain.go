package db

import (
	"time"

	"gorm.io/gorm"
)

type WatchedTrain struct {
	gorm.Model
	TrainID    string
	WatchUntil time.Time
}
