package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/d3n972/mavint/models"
	"github.com/gin-gonic/gin"
)

type TimetableController struct {
}

func (tt *TimetableController) callApi(ctx *gin.Context) []byte {
	type Payload struct {
		Type              string `json:"type"`
		TravelDate        string `json:"travelDate"`
		StationNumberCode string `json:"stationNumberCode"`
		MinCount          string `json:"minCount"`
		MaxCount          string `json:"maxCount"`
	}
	fromDate := time.Now()
	localTZ, _ := time.LoadLocation("Europe/Budapest")
	data := Payload{
		MaxCount:          "9999999",
		MinCount:          "0",
		StationNumberCode: ctx.Params.ByName("station_code"),
		TravelDate:        time.Date(fromDate.Year(), fromDate.Month(), fromDate.Day(), 0, 0, 0, 0, localTZ).Local().Format(time.RFC3339),
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
	req.Header.Set("Language", "hu")
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
	resp := tt.callApi(ctx)
	inst := models.StationTimeTable{}
	json.Unmarshal(resp, &inst)
	ctx.HTML(http.StatusOK, "timetable/tt_index", gin.H{
		"station":   inst.StationSchedulerDetails,
		"arrival":   inst.StationSchedulerDetails.ArrivalScheduler,
		"departure": inst.StationSchedulerDetails.DepartureScheduler,
	})
}
