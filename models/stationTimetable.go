package models

import (
	"fmt"
	"time"
)

var emptyStr string

type IScheduler interface {
	GetName() *string
	GetFullShortType() string
	GetCode() string
	GetIconCharacters() string
}
type STT_ArrivalScheduler struct {
	IScheduler
	AggregatedServiceIds        []interface{} `json:"aggregatedServiceIds"`
	Name                        *string       `json:"name"`
	SeatReservationCode         string        `json:"seatReservationCode"`
	Code                        string        `json:"code"`
	CompanyCode                 interface{}   `json:"companyCode"`
	Route                       interface{}   `json:"route"`
	StartStationReservationCode interface{}   `json:"startStationReservationCode"`
	EndStationReservationCode   interface{}   `json:"endStationReservationCode"`
	StartStation                struct {
		ID                            int    `json:"id"`
		IsAlias                       bool   `json:"isAlias"`
		Name                          string `json:"name"`
		Code                          string `json:"code"`
		BaseCode                      string `json:"baseCode"`
		IsInternational               bool   `json:"isInternational"`
		CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
		CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
		Country                       string `json:"country"`
		CoutryIso                     string `json:"coutryIso"`
		IsIn1081                      bool   `json:"isIn108_1"`
	} `json:"startStation"`
	EndStation struct {
		ID                            int    `json:"id"`
		IsAlias                       bool   `json:"isAlias"`
		Name                          string `json:"name"`
		Code                          string `json:"code"`
		BaseCode                      string `json:"baseCode"`
		IsInternational               bool   `json:"isInternational"`
		CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
		CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
		Country                       string `json:"country"`
		CoutryIso                     string `json:"coutryIso"`
		IsIn1081                      bool   `json:"isIn108_1"`
	} `json:"endStation"`
	StartDate        time.Time `json:"startDate"`
	OrigStartStation string    `json:"origStartStation"`
	OrigEndStation   string    `json:"origEndStation"`
	Start            time.Time `json:"start"`
	VirtualStart     bool      `json:"virtualStart"`
	Arrive           time.Time `json:"arrive"`
	VirtualArrive    bool      `json:"virtualArrive"`
	Distance         float64   `json:"distance"`
	ClosedTrackway   bool      `json:"closedTrackway"`
	FullName         string    `json:"fullName"`
	FullNameAndType  string    `json:"fullNameAndType"`
	Kinds            []struct {
		Name                string `json:"name"`
		SortName            string `json:"sortName"`
		Code                string `json:"code"`
		Priority            int    `json:"priority"`
		BackgrouColorCode   string `json:"backgrouColorCode"`
		ForegroundColorCode string `json:"foregroundColorCode"`
		Sign                struct {
			FontName  interface{} `json:"fontName"`
			Character interface{} `json:"character"`
		} `json:"sign"`
		StartStation struct {
			ID                            int    `json:"id"`
			IsAlias                       bool   `json:"isAlias"`
			Name                          string `json:"name"`
			Code                          string `json:"code"`
			BaseCode                      string `json:"baseCode"`
			IsInternational               bool   `json:"isInternational"`
			CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
			CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
			Country                       string `json:"country"`
			CoutryIso                     string `json:"coutryIso"`
			IsIn1081                      bool   `json:"isIn108_1"`
		} `json:"startStation"`
		EndStation struct {
			ID                            int    `json:"id"`
			IsAlias                       bool   `json:"isAlias"`
			Name                          string `json:"name"`
			Code                          string `json:"code"`
			BaseCode                      string `json:"baseCode"`
			IsInternational               bool   `json:"isInternational"`
			CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
			CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
			Country                       string `json:"country"`
			CoutryIso                     string `json:"coutryIso"`
			IsIn1081                      bool   `json:"isIn108_1"`
		} `json:"endStation"`
	} `json:"kinds"`
	KindsToDisplay []struct {
		Name                string `json:"name"`
		SortName            string `json:"sortName"`
		Code                string `json:"code"`
		Priority            int    `json:"priority"`
		BackgrouColorCode   string `json:"backgrouColorCode"`
		ForegroundColorCode string `json:"foregroundColorCode"`
		Sign                struct {
			FontName  interface{} `json:"fontName"`
			Character interface{} `json:"character"`
		} `json:"sign"`
		StartStation struct {
			ID                            int    `json:"id"`
			IsAlias                       bool   `json:"isAlias"`
			Name                          string `json:"name"`
			Code                          string `json:"code"`
			BaseCode                      string `json:"baseCode"`
			IsInternational               bool   `json:"isInternational"`
			CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
			CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
			Country                       string `json:"country"`
			CoutryIso                     string `json:"coutryIso"`
			IsIn1081                      bool   `json:"isIn108_1"`
		} `json:"startStation"`
		EndStation struct {
			ID                            int    `json:"id"`
			IsAlias                       bool   `json:"isAlias"`
			Name                          string `json:"name"`
			Code                          string `json:"code"`
			BaseCode                      string `json:"baseCode"`
			IsInternational               bool   `json:"isInternational"`
			CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
			CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
			Country                       string `json:"country"`
			CoutryIso                     string `json:"coutryIso"`
			IsIn1081                      bool   `json:"isIn108_1"`
		} `json:"endStation"`
	} `json:"kindsToDisplay"`
	Kind struct {
		Name                string `json:"name"`
		SortName            string `json:"sortName"`
		Code                string `json:"code"`
		Priority            int    `json:"priority"`
		BackgrouColorCode   string `json:"backgroundColorCode"`
		ForegroundColorCode string `json:"foregroundColorCode"`
		Sign                struct {
			FontName  interface{} `json:"fontName"`
			Character interface{} `json:"character"`
		} `json:"sign"`

		StartStation struct {
			ID                            int    `json:"id"`
			IsAlias                       bool   `json:"isAlias"`
			Name                          string `json:"name"`
			Code                          string `json:"code"`
			BaseCode                      string `json:"baseCode"`
			IsInternational               bool   `json:"isInternational"`
			CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
			CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
			Country                       string `json:"country"`
			CoutryIso                     string `json:"coutryIso"`
			IsIn1081                      bool   `json:"isIn108_1"`
		} `json:"startStation"`
		EndStation struct {
			ID                            int    `json:"id"`
			IsAlias                       bool   `json:"isAlias"`
			Name                          string `json:"name"`
			Code                          string `json:"code"`
			BaseCode                      string `json:"baseCode"`
			IsInternational               bool   `json:"isInternational"`
			CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
			CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
			Country                       string `json:"country"`
			CoutryIso                     string `json:"coutryIso"`
			IsIn1081                      bool   `json:"isIn108_1"`
		} `json:"endStation"`
	} `json:"kind"`
	Services []struct {
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
	ActualOrEstimatedStart  *time.Time `json:"actualOrEstimatedStart"`
	ActualOrEstimatedArrive *time.Time `json:"actualOrEstimatedArrive"`
	HavarianInfok           struct {
		AktualisKeses float64 `json:"aktualisKeses"`
		KesesiOk      string  `json:"kesesiOk"`
		HavariaInfo   string  `json:"havariaInfo"`
		UzletiInfo    string  `json:"uzletiInfo"`
		KesesInfo     string  `json:"kesesInfo"`
	} `json:"havarianInfok"`
	DirectTrains         interface{} `json:"directTrains"`
	CarrierTrains        interface{} `json:"carrierTrains"`
	StartTrack           interface{} `json:"startTrack"`
	EndTrack             interface{} `json:"endTrack"`
	JeEszkozAlapID       float64     `json:"jeEszkozAlapId"`
	FullType             string      `json:"fullType"`
	FullShortType        string      `json:"fullShortType"`
	FullNameAndPiktogram struct {
		Collection string `json:"(Collection)"`
	} `json:"fullNameAndPiktogram"`
	Footer         interface{} `json:"footer"`
	ViszonylatiJel struct {
		PiktogramFullName interface{} `json:"piktogramFullName"`
		FontSzinKod       string      `json:"fontSzinKod"`
		HatterSzinKod     string      `json:"hatterSzinKod"`
		Jel               string      `json:"jel"`
		Sign              struct {
			Character string `json:"character"`
			FontName  string `json:"fontName"`
		} `json:"sign"`
	} `json:"viszonylatiJel"`
	ViszonylatObject struct {
		StartStationCode  string      `json:"startStationCode"`
		StartTime         time.Time   `json:"startTime"`
		StartTimeZone     string      `json:"startTimeZone"`
		EndStationCode    string      `json:"endStationCode"`
		EndTime           time.Time   `json:"endTime"`
		EndTimeZone       string      `json:"endTimeZone"`
		TravelTime        float64     `json:"travelTime"`
		StartTrack        interface{} `json:"startTrack"`
		EndTrack          interface{} `json:"endTrack"`
		InnerStationCodes []string    `json:"innerStationCodes"`
	} `json:"viszonylatObject"`
	Description    interface{} `json:"description"`
	SameCar        bool        `json:"sameCar"`
	StartTimeZone  interface{} `json:"startTimeZone"`
	ArriveTimeZone string      `json:"arriveTimeZone"`
	TrainID        string      `json:"trainId"`
}

func (sch STT_ArrivalScheduler) GetName() *string {
	emptyStr = ""
	if sch.Name != nil {
		return sch.Name
	}
	return &emptyStr
}
func (sch STT_ArrivalScheduler) GetFullShortType() string {
	return sch.FullType

}
func (sch STT_ArrivalScheduler) GetCode() string {
	return sch.Code
}
func (sch STT_ArrivalScheduler) GetIconCharacters() string {
	var s string
	for _, v := range sch.Services {
		s += fmt.Sprintf("%s ", v.Sign.Character)
	}
	return s
}

type STT_DepartureScheduler struct {
	AggregatedServiceIds        []interface{} `json:"aggregatedServiceIds"`
	Name                        *string       `json:"name"`
	SeatReservationCode         string        `json:"seatReservationCode"`
	Code                        string        `json:"code"`
	CompanyCode                 interface{}   `json:"companyCode"`
	Route                       interface{}   `json:"route"`
	StartStationReservationCode interface{}   `json:"startStationReservationCode"`
	EndStationReservationCode   interface{}   `json:"endStationReservationCode"`
	StartStation                struct {
		ID                            int    `json:"id"`
		IsAlias                       bool   `json:"isAlias"`
		Name                          string `json:"name"`
		Code                          string `json:"code"`
		BaseCode                      string `json:"baseCode"`
		IsInternational               bool   `json:"isInternational"`
		CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
		CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
		Country                       string `json:"country"`
		CoutryIso                     string `json:"coutryIso"`
		IsIn1081                      bool   `json:"isIn108_1"`
	} `json:"startStation"`
	EndStation struct {
		ID                            int    `json:"id"`
		IsAlias                       bool   `json:"isAlias"`
		Name                          string `json:"name"`
		Code                          string `json:"code"`
		BaseCode                      string `json:"baseCode"`
		IsInternational               bool   `json:"isInternational"`
		CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
		CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
		Country                       string `json:"country"`
		CoutryIso                     string `json:"coutryIso"`
		IsIn1081                      bool   `json:"isIn108_1"`
	} `json:"endStation"`
	StartDate        *time.Time  `json:"startDate"`
	OrigStartStation interface{} `json:"origStartStation"`
	OrigEndStation   interface{} `json:"origEndStation"`
	Start            *time.Time  `json:"start"`
	VirtualStart     bool        `json:"virtualStart"`
	Arrive           *time.Time  `json:"arrive"`
	VirtualArrive    bool        `json:"virtualArrive"`
	Distance         float64     `json:"distance"`
	ClosedTrackway   bool        `json:"closedTrackway"`
	FullName         string      `json:"fullName"`
	FullNameAndType  string      `json:"fullNameAndType"`
	Kinds            []struct {
		Name                string `json:"name"`
		SortName            string `json:"sortName"`
		Code                string `json:"code"`
		Priority            int    `json:"priority"`
		BackgrouColorCode   string `json:"backgrouColorCode"`
		ForegroundColorCode string `json:"foregroundColorCode"`
		Sign                struct {
			FontName  interface{} `json:"fontName"`
			Character interface{} `json:"character"`
		} `json:"sign"`
		StartStation struct {
			ID                            int    `json:"id"`
			IsAlias                       bool   `json:"isAlias"`
			Name                          string `json:"name"`
			Code                          string `json:"code"`
			BaseCode                      string `json:"baseCode"`
			IsInternational               bool   `json:"isInternational"`
			CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
			CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
			Country                       string `json:"country"`
			CoutryIso                     string `json:"coutryIso"`
			IsIn1081                      bool   `json:"isIn108_1"`
		} `json:"startStation"`
		EndStation struct {
			ID                            int    `json:"id"`
			IsAlias                       bool   `json:"isAlias"`
			Name                          string `json:"name"`
			Code                          string `json:"code"`
			BaseCode                      string `json:"baseCode"`
			IsInternational               bool   `json:"isInternational"`
			CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
			CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
			Country                       string `json:"country"`
			CoutryIso                     string `json:"coutryIso"`
			IsIn1081                      bool   `json:"isIn108_1"`
		} `json:"endStation"`
	} `json:"kinds"`
	KindsToDisplay []struct {
		Name                string `json:"name"`
		SortName            string `json:"sortName"`
		Code                string `json:"code"`
		Priority            int    `json:"priority"`
		BackgrouColorCode   string `json:"backgrouColorCode"`
		ForegroundColorCode string `json:"foregroundColorCode"`
		Sign                struct {
			FontName  interface{} `json:"fontName"`
			Character interface{} `json:"character"`
		} `json:"sign"`
		StartStation struct {
			ID                            int    `json:"id"`
			IsAlias                       bool   `json:"isAlias"`
			Name                          string `json:"name"`
			Code                          string `json:"code"`
			BaseCode                      string `json:"baseCode"`
			IsInternational               bool   `json:"isInternational"`
			CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
			CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
			Country                       string `json:"country"`
			CoutryIso                     string `json:"coutryIso"`
			IsIn1081                      bool   `json:"isIn108_1"`
		} `json:"startStation"`
		EndStation struct {
			ID                            int    `json:"id"`
			IsAlias                       bool   `json:"isAlias"`
			Name                          string `json:"name"`
			Code                          string `json:"code"`
			BaseCode                      string `json:"baseCode"`
			IsInternational               bool   `json:"isInternational"`
			CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
			CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
			Country                       string `json:"country"`
			CoutryIso                     string `json:"coutryIso"`
			IsIn1081                      bool   `json:"isIn108_1"`
		} `json:"endStation"`
	} `json:"kindsToDisplay"`
	Kind struct {
		Name                string `json:"name"`
		SortName            string `json:"sortName"`
		Code                string `json:"code"`
		Priority            int    `json:"priority"`
		BackgrouColorCode   string `json:"backgroundColorCode"`
		ForegroundColorCode string `json:"foregroundColorCode"`
		Sign                struct {
			FontName  interface{} `json:"fontName"`
			Character interface{} `json:"character"`
		} `json:"sign"`
		StartStation struct {
			ID                            int    `json:"id"`
			IsAlias                       bool   `json:"isAlias"`
			Name                          string `json:"name"`
			Code                          string `json:"code"`
			BaseCode                      string `json:"baseCode"`
			IsInternational               bool   `json:"isInternational"`
			CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
			CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
			Country                       string `json:"country"`
			CoutryIso                     string `json:"coutryIso"`
			IsIn1081                      bool   `json:"isIn108_1"`
		} `json:"startStation"`
		EndStation struct {
			ID                            int    `json:"id"`
			IsAlias                       bool   `json:"isAlias"`
			Name                          string `json:"name"`
			Code                          string `json:"code"`
			BaseCode                      string `json:"baseCode"`
			IsInternational               bool   `json:"isInternational"`
			CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
			CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
			Country                       string `json:"country"`
			CoutryIso                     string `json:"coutryIso"`
			IsIn1081                      bool   `json:"isIn108_1"`
		} `json:"endStation"`
	} `json:"kind"`
	Services []struct {
		ListOrder                   interface{} `json:"listOrder"`
		Description                 string      `json:"description"`
		RestrictiveStartStationCode interface{} `json:"restrictiveStartStationCode"`
		RestrictiveEndStationCode   interface{} `json:"restrictiveEndStationCode"`
		Sign                        struct {
			FontName  string `json:"fontName"`
			Character string `json:"character"`
		} `json:"sign"`
		TrainStopKind interface{} `json:"trainStopKind"`
	} `json:"services"`
	ActualOrEstimatedStart  *time.Time `json:"actualOrEstimatedStart"`
	ActualOrEstimatedArrive *time.Time `json:"actualOrEstimatedArrive"`
	HavarianInfok           struct {
		AktualisKeses float64     `json:"aktualisKeses"`
		KesesiOk      interface{} `json:"kesesiOk"`
		HavariaInfo   interface{} `json:"havariaInfo"`
		UzletiInfo    interface{} `json:"uzletiInfo"`
		KesesInfo     string      `json:"kesesInfo"`
	} `json:"havarianInfok"`
	DirectTrains         interface{} `json:"directTrains"`
	CarrierTrains        interface{} `json:"carrierTrains"`
	StartTrack           string      `json:"startTrack"`
	EndTrack             string      `json:"endTrack"`
	JeEszkozAlapID       float64     `json:"jeEszkozAlapId"`
	FullType             string      `json:"fullType"`
	FullShortType        string      `json:"fullShortType"`
	FullNameAndPiktogram struct {
		Collection string `json:"(Collection)"`
	} `json:"fullNameAndPiktogram"`
	Footer         string `json:"footer"`
	ViszonylatiJel struct {
		PiktogramFullName string `json:"piktogramFullName"`
		FontSzinKod       string `json:"fontSzinKod"`
		HatterSzinKod     string `json:"hatterSzinKod"`
		Jel               string `json:"jel"`
		Sign              struct {
			Character string `json:"character"`
			FontName  string `json:"fontName"`
		} `json:"sign"`
	} `json:"viszonylatiJel"`
	ViszonylatObject struct {
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
	} `json:"viszonylatObject"`
	Description    interface{} `json:"description"`
	SameCar        bool        `json:"sameCar"`
	StartTimeZone  string      `json:"startTimeZone"`
	ArriveTimeZone interface{} `json:"arriveTimeZone"`
	TrainID        string      `json:"trainId"`
}

func (sch STT_DepartureScheduler) GetName() *string {
	emptyStr = ""
	if sch.Name != nil {
		return sch.Name
	}
	return &emptyStr
}
func (sch STT_DepartureScheduler) GetFullShortType() string {
	return sch.FullType
}
func (sch STT_DepartureScheduler) GetCode() string {
	return sch.Code
}
func (sch STT_DepartureScheduler) GetIconCharacters() string {
	var s string
	for _, v := range sch.Services {
		s += fmt.Sprintf("%s ", v.Sign.Character)
	}
	return s
}

type StationTimeTable struct {
	TrainSchedulerDetails   interface{} `json:"trainSchedulerDetails"`
	StationSchedulerDetails struct {
		Station struct {
			ID                            int         `json:"id"`
			IsAlias                       bool        `json:"isAlias"`
			Name                          string      `json:"name"`
			Code                          string      `json:"code"`
			BaseCode                      interface{} `json:"baseCode"`
			IsInternational               bool        `json:"isInternational"`
			CanUseForOfferRequest         bool        `json:"canUseForOfferRequest"`
			CanUseForPessengerInformation bool        `json:"canUseForPessengerInformation"`
			Country                       interface{} `json:"country"`
			CoutryIso                     interface{} `json:"coutryIso"`
			IsIn1081                      bool        `json:"isIn108_1"`
		} `json:"station"`
		ArrivalScheduler   []STT_ArrivalScheduler   `json:"arrivalScheduler"`
		DepartureScheduler []STT_DepartureScheduler `json:"departureScheduler"`
		Services           []struct {
			ListOrder                   interface{} `json:"listOrder"`
			Description                 string      `json:"description"`
			RestrictiveStartStationCode interface{} `json:"restrictiveStartStationCode"`
			RestrictiveEndStationCode   interface{} `json:"restrictiveEndStationCode"`
			Sign                        struct {
				FontName  string `json:"fontName"`
				Character string `json:"character"`
			} `json:"sign"`
			TrainStopKind interface{} `json:"trainStopKind"`
		} `json:"services"`
		MoreResult bool `json:"moreResult"`
	} `json:"stationSchedulerDetails"`
	RouteSchedulerDetails interface{} `json:"routeSchedulerDetails"`
}
