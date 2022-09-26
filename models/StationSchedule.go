package models

type StationSchedule struct {
	TrainSchedulerDetails   []TrainSchedulerDetails `json:"trainSchedulerDetails,omitempty"`
	StationSchedulerDetails StationSchedulerDetails `json:"stationSchedulerDetails,omitempty"`
	RouteSchedulerDetails   []RouteSchedulerDetails `json:"routeSchedulerDetails,omitempty"`
}
