package db

import "gorm.io/gorm"

type EngineWorkday struct {
	gorm.Model
	UIC         string
	Date        string
	JobType     string
	TrainNumber *string
}
