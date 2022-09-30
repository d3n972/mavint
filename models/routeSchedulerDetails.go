package models

import "time"

type RouteSchedulerDetails struct {
	Trains                     []Train                      `json:"trains"`
	FullDistance               float64                      `json:"fullDistance"`
	StartTime                  time.Time                    `json:"startTime"`
	EndTime                    time.Time                    `json:"endTime"`
	TravelTime                 string                       `json:"travelTime"`
	TransfersCount             int                          `json:"transfersCount"`
	StartStation               StartStation                 `json:"startStation"`
	EndStation                 EndStation                   `json:"endStation"`
	TrainsFullName             string                       `json:"trainsFullName"`
	TrainsFullNameAndPiktogram []TrainsFullNameAndPiktogram `json:"trainsFullNameAndPiktogram"`
}
