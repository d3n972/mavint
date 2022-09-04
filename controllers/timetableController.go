package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TimetableController struct {
}

func (tt *TimetableController) callApi() []byte {
	type Payload struct {
		Type              string `json:"type"`
		TravelDate        string `json:"travelDate"`
		StationNumberCode string `json:"stationNumberCode"`
		MinCount          string `json:"minCount"`
		MaxCount          string `json:"maxCount"`
	}

	data := Payload{
		MaxCount:          "9999999",
		MinCount:          "0",
		StationNumberCode: "005510033",
		TravelDate:        "2022-09-03T00:00:00.000Z", //time.Now().Local().Format("2006-01-02T03:04:05.000Z"),
		Type:              "StationInfo",
	}
	fmt.Printf("data.TravelDate: %v\n", data.TravelDate)
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://jegy-a.mav.hu/IK_API_PROD/api/InformationApi/GetTimetable", body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en-GB;q=0.9,en;q=0.8,hu-HU;q=0.7,hu;q=0.6")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Language", "en")
	req.Header.Set("Origin", "https://jegy.mav.hu")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://jegy.mav.hu/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	req.Header.Set("Usersessionid", "\"b0d2c001-2b7a-41fa-9488-49d48f295b8d\"")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"104\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"104\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")

	resp, err := http.DefaultClient.Do(req)
	ret, _ := io.ReadAll(resp.Body)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	return ret
}
func (tt *TimetableController) Render(ctx *gin.Context) {
	resp := tt.callApi()
	inst := StationTimeTable{}
	json.Unmarshal(resp, &inst)
	ctx.HTML(http.StatusOK, "timetable/tt_index", gin.H{
		"w": inst.StationSchedulerDetails.ArrivalScheduler,
	})
}

type STT_ArrivalScheduler struct {
	AggregatedServiceIds        []interface{} `json:"aggregatedServiceIds"`
	Name                        interface{}   `json:"name"`
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
		AktualisKeses int    `json:"aktualisKeses"`
		KesesiOk      string `json:"kesesiOk"`
		HavariaInfo   string `json:"havariaInfo"`
		UzletiInfo    string `json:"uzletiInfo"`
		KesesInfo     string `json:"kesesInfo"`
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
type STT_DepartureScheduler struct {
	AggregatedServiceIds        []interface{} `json:"aggregatedServiceIds"`
	Name                        interface{}   `json:"name"`
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
	ActualOrEstimatedStart  interface{} `json:"actualOrEstimatedStart"`
	ActualOrEstimatedArrive interface{} `json:"actualOrEstimatedArrive"`
	HavarianInfok           struct {
		AktualisKeses float64     `json:"aktualisKeses"`
		KesesiOk      interface{} `json:"kesesiOk"`
		HavariaInfo   interface{} `json:"havariaInfo"`
		UzletiInfo    interface{} `json:"uzletiInfo"`
		KesesInfo     string      `json:"kesesInfo"`
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
	StartTimeZone  string      `json:"startTimeZone"`
	ArriveTimeZone interface{} `json:"arriveTimeZone"`
	TrainID        string      `json:"trainId"`
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
