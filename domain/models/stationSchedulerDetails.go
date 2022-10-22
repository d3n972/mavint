package models

type StationSchedulerDetails struct {
	Station            Station              `json:"station"`
	ArrivalScheduler   []ArrivalScheduler   `json:"arrivalScheduler"`
	DepartureScheduler []DepartureScheduler `json:"departureScheduler"`
	Services           []Services           `json:"services"`
	MoreResult         bool                 `json:"moreResult"`
}
