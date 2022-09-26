package models

import (
	"time"
)

type Scheduler struct {
	Station                 Station     `json:"station"`
	Arrive                  *time.Time  `json:"arrive"`
	Start                   *time.Time  `json:"start"`
	ActualOrEstimatedArrive *time.Time  `json:"actualOrEstimatedArrive"`
	ActualOrEstimatedStart  *time.Time  `json:"actualOrEstimatedStart"`
	StartTrack              string      `json:"startTrack"`
	EndTrack                string      `json:"endTrack"`
	Services                []Services  `json:"services"`
	StopKind                int         `json:"stopKind"`
	StopService             StopService `json:"stopService"`
	Distance                int         `json:"distance"`
	StartTimeZone           string      `json:"startTimeZone"`
	ArriveTimeZone          string      `json:"arriveTimeZone"`
}
