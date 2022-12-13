package dao

import (
	"time"

	"github.com/d3n972/mavint/domain"
	"github.com/d3n972/mavint/domain/repository"
)

type WatchedTrainDAO struct {
	TrainID    string
	WatchUntil time.Time
}

func (e2 WatchedTrainDAO) TableName() string {
	return "watched_trains"
}
func (e2 WatchedTrainDAO) FromEntity(e domain.WatchedTrain) repository.DAO[domain.WatchedTrain] {
	return WatchedTrainDAO{
		TrainID:    e.TrainID,
		WatchUntil: e.WatchUntil,
	}
}

func (e2 WatchedTrainDAO) ToEntity() domain.WatchedTrain {
	return domain.WatchedTrain{
		TrainID:    e2.TrainID,
		WatchUntil: e2.WatchUntil,
	}
}
