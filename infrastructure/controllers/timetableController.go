package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-redis/redis/v9"

	"github.com/d3n972/mavint/domain/models"
	"github.com/gin-gonic/gin"
)

type TimetableController struct {
}

func (tt *TimetableController) callApi(ctx *gin.Context) []byte {
	cacheId := "timetable:" + ctx.Params.ByName("station_code")
	hRdb, _ := ctx.Get("cache")
	rdb := hRdb.(*redis.Client)
	var responseBytes []byte
	type Payload struct {
		Type              string `json:"type"`
		TravelDate        string `json:"travelDate"`
		StationNumberCode string `json:"stationNumberCode"`
		MinCount          string `json:"minCount"`
		MaxCount          string `json:"maxCount"`
	}
	if exists := rdb.Exists(context.Background(), cacheId).Val(); exists == 0 {
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
		fmt.Printf("Saving timetable for key: %s\n", cacheId)
		rdb.Set(context.Background(), cacheId, ret, 5*time.Minute)
		responseBytes, _ = rdb.Get(context.Background(), cacheId).Bytes()

	} else {
		responseBytes, _ = rdb.Get(context.Background(), cacheId).Bytes()
	}
	return responseBytes
}
func (tt *TimetableController) getStations(ctx *gin.Context) []byte {
	req, err := http.NewRequest("POST", "https://jegy-a.mav.hu/IK_API_PROD/api/OfferRequestApi/GetStationList", nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en-GB;q=0.9,en;q=0.8,hu-HU;q=0.7,hu;q=0.6")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "0")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Language", "hu")
	req.Header.Set("Origin", "https://jegy.mav.hu")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://jegy.mav.hu/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	req.Header.Set("Usersessionid", "\"0ccc3dc1-8365-474f-99fe-e433c107df19\"")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"105\", \"Not)A;Brand\";v=\"8\", \"Chromium\";v=\"105\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return b
}
func (tt *TimetableController) Render(ctx *gin.Context) {
	hc := models.HavariaCache{}
	resp := tt.callApi(ctx)
	R, _ := ctx.Get("cache")
	inst := models.StationTimeTable{}
	json.Unmarshal(resp, &inst)
	hcb := []byte(R.(*redis.Client).Get(context.TODO(), "havariaCache").Val())
	json.Unmarshal(hcb, &hc)

	ctx.HTML(http.StatusOK, "timetable/tt_next", gin.H{
		"delays":    hc,
		"station":   inst.StationSchedulerDetails,
		"arrival":   inst.StationSchedulerDetails.ArrivalScheduler,
		"departure": inst.StationSchedulerDetails.DepartureScheduler,
	})
}
func (tt *TimetableController) RenderSelectorPage(ctx *gin.Context) {
	resp := tt.getStations(ctx)
	m := []models.StationsResponse{}
	json.Unmarshal(resp, &m)
	ctx.HTML(http.StatusOK, "timetable/station_selector", gin.H{
		"stations": m,
	})
}
