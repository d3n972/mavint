package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TrainDetailsController struct {
}
type Payload struct {
	Type        string `json:"type"`
	TravelDate  string `json:"travelDate"`
	MinCount    string `json:"minCount"`
	MaxCount    string `json:"maxCount"`
	TrainID     int    `json:"trainId,omitempty"`
	TrainNumber string `json:"trainNumber,omitempty"`
}

func (c *TrainDetailsController) getApiResponse(ctx *gin.Context) []byte {
	var data Payload
	if ctx.Params.ByName("train_id") != "" {
		i, _ := strconv.Atoi(ctx.Params.ByName("train_id"))
		data = Payload{
			MaxCount:   "9999999",
			MinCount:   "0",
			TrainID:    i,
			TravelDate: "2022-08-31T22:00:00.000Z",
			Type:       "TrainInfo",
		}
	} else {
		data = Payload{
			MaxCount:    "9999999",
			MinCount:    "0",
			TrainNumber: "707",
			TravelDate:  "2022-08-31T22:00:00.000Z",
			Type:        "TrainInfo",
		}
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://jegy-a.mav.hu/IK_API_PROD/api/InformationApi/GetTimetable", body)
	if err != nil {
		// handle err
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
	req.Header.Set("Usersessionid", "\"4d793891-0e70-45b5-a5c6-64eddbba2532\"")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"104\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"104\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	respBytes, _ := io.ReadAll(resp.Body)
	return respBytes
}
func (c *TrainDetailsController) Render(ctx *gin.Context) {
	type TrainDetailsResponse struct {
		TrainSchedulerDetails []struct {
			Train struct {
				AggregatedServiceIds        []interface{} `json:"aggregatedServiceIds"`
				Name                        string        `json:"name"`
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
				StartDate        interface{} `json:"startDate"`
				OrigStartStation interface{} `json:"origStartStation"`
				OrigEndStation   interface{} `json:"origEndStation"`
				Start            time.Time   `json:"start"`
				VirtualStart     bool        `json:"virtualStart"`
				Arrive           time.Time   `json:"arrive"`
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
						FontName  string `json:"fontName"`
						Character string `json:"character"`
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
						FontName  string `json:"fontName"`
						Character string `json:"character"`
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
						FontName  string `json:"fontName"`
						Character string `json:"character"`
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
					ListOrder                   string      `json:"listOrder"`
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
				DirectTrains []struct {
					Train struct {
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
						StartDate               interface{}   `json:"startDate"`
						OrigStartStation        interface{}   `json:"origStartStation"`
						OrigEndStation          interface{}   `json:"origEndStation"`
						Start                   interface{}   `json:"start"`
						VirtualStart            bool          `json:"virtualStart"`
						Arrive                  interface{}   `json:"arrive"`
						VirtualArrive           bool          `json:"virtualArrive"`
						Distance                float64       `json:"distance"`
						ClosedTrackway          bool          `json:"closedTrackway"`
						FullName                string        `json:"fullName"`
						FullNameAndType         string        `json:"fullNameAndType"`
						Kinds                   interface{}   `json:"kinds"`
						KindsToDisplay          interface{}   `json:"kindsToDisplay"`
						Kind                    interface{}   `json:"kind"`
						Services                []interface{} `json:"services"`
						ActualOrEstimatedStart  interface{}   `json:"actualOrEstimatedStart"`
						ActualOrEstimatedArrive interface{}   `json:"actualOrEstimatedArrive"`
						HavarianInfok           interface{}   `json:"havarianInfok"`
						DirectTrains            interface{}   `json:"directTrains"`
						CarrierTrains           interface{}   `json:"carrierTrains"`
						StartTrack              interface{}   `json:"startTrack"`
						EndTrack                interface{}   `json:"endTrack"`
						JeEszkozAlapID          float64       `json:"jeEszkozAlapId"`
						FullType                string        `json:"fullType"`
						FullShortType           string        `json:"fullShortType"`
						FullNameAndPiktogram    struct {
							Collection string `json:"(Collection)"`
						} `json:"fullNameAndPiktogram"`
						Footer           interface{} `json:"footer"`
						ViszonylatiJel   interface{} `json:"viszonylatiJel"`
						ViszonylatObject interface{} `json:"viszonylatObject"`
						Description      interface{} `json:"description"`
						SameCar          bool        `json:"sameCar"`
						StartTimeZone    interface{} `json:"startTimeZone"`
						ArriveTimeZone   interface{} `json:"arriveTimeZone"`
						TrainID          string      `json:"trainId"`
					} `json:"train"`
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
					Footer interface{} `json:"footer"`
				} `json:"directTrains"`
				CarrierTrains        interface{} `json:"carrierTrains"`
				StartTrack           interface{} `json:"startTrack"`
				EndTrack             interface{} `json:"endTrack"`
				JeEszkozAlapID       float64     `json:"jeEszkozAlapId"`
				FullType             string      `json:"fullType"`
				FullShortType        string      `json:"fullShortType"`
				FullNameAndPiktogram struct {
					Collection string `json:"(Collection)"`
				} `json:"fullNameAndPiktogram"`
				Footer           string      `json:"footer"`
				ViszonylatiJel   interface{} `json:"viszonylatiJel"`
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
				ArriveTimeZone interface{} `json:"arriveTimeZone"`
				TrainID        string      `json:"trainId"`
			} `json:"train"`
			Scheduler []struct {
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
					CoutryIso                     string `json:"coutryIso"`
					IsIn1081                      bool   `json:"isIn108_1"`
				} `json:"station"`
				Arrive                  interface{} `json:"arrive"`
				Start                   time.Time   `json:"start"`
				ActualOrEstimatedArrive interface{} `json:"actualOrEstimatedArrive"`
				ActualOrEstimatedStart  time.Time   `json:"actualOrEstimatedStart"`
				StartTrack              interface{} `json:"startTrack"`
				EndTrack                interface{} `json:"endTrack"`
				Services                []struct {
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
				StopKind    int `json:"stopKind"`
				StopService struct {
					ListOrder                   interface{} `json:"listOrder"`
					Description                 interface{} `json:"description"`
					RestrictiveStartStationCode interface{} `json:"restrictiveStartStationCode"`
					RestrictiveEndStationCode   interface{} `json:"restrictiveEndStationCode"`
					Sign                        interface{} `json:"sign"`
					TrainStopKind               interface{} `json:"trainStopKind"`
				} `json:"stopService"`
				Distance       int         `json:"distance"`
				StartTimeZone  string      `json:"startTimeZone"`
				ArriveTimeZone interface{} `json:"arriveTimeZone"`
			} `json:"scheduler"`
		} `json:"trainSchedulerDetails"`
		StationSchedulerDetails interface{} `json:"stationSchedulerDetails"`
		RouteSchedulerDetails   interface{} `json:"routeSchedulerDetails"`
	}

	apiresp := c.getApiResponse(ctx)
	instance := TrainDetailsResponse{}
	json.Unmarshal(apiresp, &instance)

	dJSON, _ := json.MarshalIndent(instance.TrainSchedulerDetails, "", "    ")

	ctx.HTML(http.StatusOK, "traininfo/info", gin.H{
		"info": instance.TrainSchedulerDetails[0],
		"data": string(dJSON),
	})
}
