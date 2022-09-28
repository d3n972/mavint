package models

import (
	"fmt"
)

var emptyStr string

type IScheduler interface {
	GetName() *string
	GetFullShortType() string
	GetCode() string
	GetIconCharacters() string
}

func (sch ArrivalScheduler) GetName() *string {
	emptyStr = ""
	if sch.Name != nil {
		return sch.Name
	}
	return &emptyStr
}
func (sch ArrivalScheduler) GetFullShortType() string {
	return sch.FullType

}
func (sch ArrivalScheduler) GetCode() string {
	return sch.Code
}
func (sch ArrivalScheduler) GetIconCharacters() string {
	var s string
	for _, v := range sch.Services {
		s += fmt.Sprintf("%s ", v.Sign.Character)
	}
	return s
}
func (sch DepartureScheduler) GetName() *string {
	emptyStr = ""
	if sch.Name != nil {
		return sch.Name
	}
	return &emptyStr
}
func (sch DepartureScheduler) GetFullShortType() string {
	return sch.FullType

}
func (sch DepartureScheduler) GetCode() string {
	return sch.Code
}
func (sch DepartureScheduler) GetIconCharacters() string {
	var s string
	for _, v := range sch.Services {
		s += fmt.Sprintf("%s ", v.Sign.Character)
	}
	return s
}

type StationTimeTable struct {
	TrainSchedulerDetails   TrainSchedulerDetails   `json:"trainSchedulerDetails"`
	StationSchedulerDetails StationSchedulerDetails `json:"stationSchedulerDetails"`
	RouteSchedulerDetails   RouteSchedulerDetails   `json:"routeSchedulerDetails"`
}
