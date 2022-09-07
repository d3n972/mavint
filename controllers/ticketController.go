package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/d3n972/mavint/models"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type TicketController struct {
}
type Passangers struct {
	PassengerCount        int    `json:"passengerCount"`
	PassengerID           int    `json:"passengerId"`
	CustomerTypeKey       string `json:"customerTypeKey"`
	CustomerDiscountsKeys []int  `json:"customerDiscountsKeys"`
}

func (t *TicketController) apiCall() {

	type Payload struct {
		Offerkind                            string       `json:"offerkind"`
		StartStationCode                     string       `json:"startStationCode"`
		InnerStationsCodes                   []string     `json:"innerStationsCodes"`
		EndStationCode                       string       `json:"endStationCode"`
		Passangers                           []Passangers `json:"passangers"`
		IsOneWayTicket                       bool         `json:"isOneWayTicket"`
		IsTravelEndTime                      bool         `json:"isTravelEndTime"`
		IsSupplementaryTicketsOnly           bool         `json:"isSupplementaryTicketsOnly"`
		TravelStartDate/*time.Time*/ string               `json:"travelStartDate"`
		TravelReturnDate/*time.Time*/ string              `json:"travelReturnDate"`
		SelectedServices                     []int        `json:"selectedServices"`
		SelectedSearchServices               []string     `json:"selectedSearchServices"`
		EszkozSzamok                         []string     `json:"eszkozSzamok"`
		IsOfDetailedSearch                   bool         `json:"isOfDetailedSearch"`
	}
	//tRet, _ := time.Parse("2006-01-02T15:04:05-07:00", "2022-09-06T00:30:00+02:00")
	//tStart, _ := time.Parse("2006-01-02T15:04:05-07:00", "2022-09-06T00:00:00+02:00")

	data := Payload{
		EszkozSzamok:               []string{},
		InnerStationsCodes:         []string{},
		IsOfDetailedSearch:         false,
		IsOneWayTicket:             true,
		IsSupplementaryTicketsOnly: false,
		IsTravelEndTime:            false,
		Offerkind:                  "1",
		Passangers: []Passangers{
			Passangers{
				PassengerCount:        1,
				PassengerID:           0,
				CustomerTypeKey:       "HU_44_026-065",
				CustomerDiscountsKeys: []int{},
			},
		},
		SelectedSearchServices: []string{
			"BUDAPESTI_HELYI_KOZLEKEDESSEL",
		},
		SelectedServices: []int{52},
		StartStationCode: "005517111",
		EndStationCode:   "005513912",
		TravelReturnDate: "2022-09-07T00:30:00+02:00",
		TravelStartDate:  "2022-09-07T00:00:00+02:00",
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://jegy-a.mav.hu/IK_API_PROD/api/OfferRequestApi/GetOfferRequest", body)
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
	req.Header.Set("Usersessionid", "\"2ad1c19d-9f6e-4193-b7d1-f9f869156d83\"")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"104\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"104\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")

	resp, err := http.DefaultClient.Do(req)
	respBytes, _ := io.ReadAll(resp.Body)
	m := models.TicketListingResponse{}
	json.Unmarshal(respBytes, &m)
	route := m.Route[1]
	ticket := route.Details.Tickets[0]
	sumPrice := 0.0
	for _, s := range route.Details.Tickets {
		sumPrice += s.GrossPrice.Amount
	}
	fmt.Printf(`
%s
%s %s
** 1 db - %s. kocsiosztály - Egy útra
Km: MÁV/0/%.1f
Érvényes:  2022.08.27. 00:00
			 2022.08.28. 23:59
--------------------------------------
Ára összesen:	%s  **%.0f

`, route.Details.Routes[0].StartStation.Name,
		route.Details.Routes[1].DestionationStation.Name,
		route.Details.Routes[0].TouchedStationsString,
		route.Details.Routes[0].TravelClasses[0].Name,
		route.Details.Routes[0].Distance,
		ticket.GrossPrice.Currency.Key,
		sumPrice,
	)

	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}
func (t *TicketController) Render(ctx *gin.Context) {
	t.apiCall()
	ctx.JSON(http.StatusOK, gin.H{"ok": true})
}
