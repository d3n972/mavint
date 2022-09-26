package models

type TrainSchedulerDetails struct {
	Train     Train       `json:"train"`
	Scheduler []Scheduler `json:"scheduler"`
}
