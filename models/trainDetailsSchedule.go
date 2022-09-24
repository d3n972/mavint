package models

import "time"

type TD_Scheduler struct {
	Station struct {
		ID                            int    `json:"id"`
		IsAlias                       bool   `json:"isAlias"`
		Name                          string `json:"name"`
		Code                          string `json:"code"`
		BaseCode                      string `json:"baseCode"`
		IsInternational               bool   `json:"isInternational"`
		CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
		CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
		Country                       string `json:"country"`
		CountryIso                    string `json:"coutryIso"`
		IsIn1081                      bool   `json:"isIn108_1"`
	} `json:"station"`
	Arrive                  *time.Time `json:"arrive"`
	Start                   *time.Time `json:"start"`
	ActualOrEstimatedArrive *time.Time `json:"actualOrEstimatedArrive"`
	ActualOrEstimatedStart  *time.Time `json:"actualOrEstimatedStart"`
	StartTrack              string     `json:"startTrack"`
	EndTrack                string     `json:"endTrack"`
	Services                []struct {
		ListOrder                   interface{} `json:"listOrder"`
		Description                 string      `json:"description"`
		RestrictiveStartStationCode interface{} `json:"restrictiveStartStationCode"`
		RestrictiveEndStationCode   interface{} `json:"restrictiveEndStationCode"`
		Sign                        struct {
			FontName  string `json:"fontName"`
			Character string `json:"character"`
		} `json:"sign"`
		TrainStopKind string `json:"trainStopKind"`
	} `json:"services"`
	StopKind    int `json:"stopKind"`
	StopService struct {
		ListOrder                   interface{} `json:"listOrder"`
		Description                 interface{} `json:"description"`
		RestrictiveStartStationCode interface{} `json:"restrictiveStartStationCode"`
		RestrictiveEndStationCode   interface{} `json:"restrictiveEndStationCode"`
		Sign                        interface{} `json:"sign"`
		TrainStopKind               string      `json:"trainStopKind"`
	} `json:"stopService"`
	Distance       int         `json:"distance"`
	StartTimeZone  string      `json:"startTimeZone"`
	ArriveTimeZone interface{} `json:"arriveTimeZone"`
}
