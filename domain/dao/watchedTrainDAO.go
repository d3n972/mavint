package dao

import (
	"github.com/d3n972/mavint/domain"
	"github.com/d3n972/mavint/domain/repository"
)

type WatchedTrainDAO struct {
	__e domain.WatchedTrain
}

func (e2 WatchedTrainDAO) FromEntity(e domain.WatchedTrain) repository.DAO[domain.WatchedTrain] {
	//TODO implement me
	panic("implement me")
}

func (e2 WatchedTrainDAO) ToEntity() domain.WatchedTrain {
	//TODO implement me
	panic("implement me")
}
