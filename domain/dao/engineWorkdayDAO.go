package dao

import (
	"github.com/d3n972/mavint/domain"
	"github.com/d3n972/mavint/domain/repository"
	"gorm.io/gorm"
)

type EngineWorkdayDAO struct {
	gorm.Model
	UIC            string
	Date           string
	JobType        string
	TrainNumber    *string
	NearestStation *string
}

func (e2 EngineWorkdayDAO) FromEntity(e domain.EngineWorkday) repository.DAO[domain.EngineWorkday] {
	//TODO implement me
	panic("implement me")
}

func (e2 EngineWorkdayDAO) ToEntity() domain.EngineWorkday {
	//TODO implement me
	panic("implement me")
}
