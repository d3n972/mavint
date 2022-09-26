package models

import "time"

type ViszonylatObject struct {
	StartStationCode  string    `json:"startStationCode"`
	StartTime         time.Time `json:"startTime"`
	StartTimeZone     string    `json:"startTimeZone"`
	EndStationCode    string    `json:"endStationCode"`
	EndTime           time.Time `json:"endTime"`
	EndTimeZone       string    `json:"endTimeZone"`
	TravelTime        float64   `json:"travelTime"`
	StartTrack        string    `json:"startTrack"`
	EndTrack          string    `json:"endTrack"`
	InnerStationCodes []string  `json:"innerStationCodes"`
}
