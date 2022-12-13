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

func (e2 EngineWorkdayDAO) TableName() string {
	return "engine_workdays"
}
func (e2 EngineWorkdayDAO) FromEntity(e domain.EngineWorkday) repository.DAO[domain.EngineWorkday] {
	return EngineWorkdayDAO{
		UIC:            e2.UIC,
		Date:           e2.Date,
		JobType:        e2.JobType,
		TrainNumber:    e2.TrainNumber,
		NearestStation: e2.NearestStation,
	}
}

func (e2 EngineWorkdayDAO) ToEntity() domain.EngineWorkday {
	return domain.EngineWorkday{
		UIC:            e2.UIC,
		Date:           e2.Date,
		JobType:        e2.JobType,
		TrainNumber:    e2.TrainNumber,
		NearestStation: e2.NearestStation,
	}
}
