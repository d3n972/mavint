package models

type DirectTrains struct {
	Train        Train        `json:"train"`
	StartStation StartStation `json:"startStation"`
	EndStation   EndStation   `json:"endStation"`
	Footer       string       `json:"footer"`
}
