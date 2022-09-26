package models

import (
	"fmt"
	"time"
)

type Train struct {
	AggregatedServiceIds        []string             `json:"aggregatedServiceIds"`
	Name                        *string              `json:"name"`
	SeatReservationCode         string               `json:"seatReservationCode"`
	Code                        string               `json:"code"`
	CompanyCode                 string               `json:"companyCode"`
	Route                       string               `json:"route"`
	StartStationReservationCode string               `json:"startStationReservationCode"`
	EndStationReservationCode   string               `json:"endStationReservationCode"`
	StartStation                StartStation         `json:"startStation"`
	EndStation                  EndStation           `json:"endStation"`
	StartDate                   time.Time            `json:"startDate"`
	OrigStartStation            Station              `json:"origStartStation"`
	OrigEndStation              Station              `json:"origEndStation"`
	Start                       time.Time            `json:"start"`
	VirtualStart                bool                 `json:"virtualStart"`
	Arrive                      time.Time            `json:"arrive"`
	VirtualArrive               bool                 `json:"virtualArrive"`
	Distance                    float64              `json:"distance"`
	ClosedTrackway              bool                 `json:"closedTrackway"`
	FullName                    string               `json:"fullName"`
	FullNameAndType             string               `json:"fullNameAndType"`
	Kinds                       []Kinds              `json:"kinds"`
	KindsToDisplay              []KindsToDisplay     `json:"kindsToDisplay"`
	Kind                        Kind                 `json:"kind"`
	Services                    []Services           `json:"services"`
	ActualOrEstimatedStart      time.Time            `json:"actualOrEstimatedStart"`
	ActualOrEstimatedArrive     time.Time            `json:"actualOrEstimatedArrive"`
	HavarianInfok               HavarianInfok        `json:"havarianInfok"`
	DirectTrains                []DirectTrains       `json:"directTrains"`
	CarrierTrains               interface{}          `json:"carrierTrains"`
	StartTrack                  string               `json:"startTrack"`
	EndTrack                    string               `json:"endTrack"`
	JeEszkozAlapID              float64              `json:"jeEszkozAlapId"`
	FullType                    string               `json:"fullType"`
	FullShortType               string               `json:"fullShortType"`
	FullNameAndPiktogram        FullNameAndPiktogram `json:"fullNameAndPiktogram"`
	Footer                      string               `json:"footer"`
	ViszonylatiJel              ViszonylatiJel       `json:"viszonylatiJel"`
	ViszonylatObject            ViszonylatObject     `json:"viszonylatObject"`
	Description                 string               `json:"description"`
	SameCar                     bool                 `json:"sameCar"`
	StartTimeZone               string               `json:"startTimeZone"`
	ArriveTimeZone              string               `json:"arriveTimeZone"`
	TrainID                     string               `json:"trainId"`
}

func (sch Train) GetName() *string {
	emptyStr = ""
	if sch.Name != nil {
		return sch.Name
	}
	return &emptyStr
}
func (sch Train) GetFullShortType() string {
	return sch.FullType

}
func (sch Train) GetCode() string {
	return sch.Code
}
func (sch Train) GetIconCharacters() string {
	var s string
	for _, v := range sch.Services {
		s += fmt.Sprintf("%s ", v.Sign.Character)
	}
	return s
}
