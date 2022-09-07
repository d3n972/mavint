package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/d3n972/mavint/models"
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
		i, _ := strconv.Atoi(ctx.Params.ByName("train"))
		data = Payload{
			MaxCount:   "9999999",
			MinCount:   "0",
			TrainID:    i,
			TravelDate: time.Now().Format("2006-01-02T00:00:00Z"),
			Type:       "TrainInfo",
		}
	} else {
		fmt.Println("xxx" + time.Now().Format("2006-01-02T00:00:00Z"))
		data = Payload{
			MaxCount:    "9999999",
			MinCount:    "0",
			TrainNumber: ctx.Query("tid"),
			TravelDate:  time.Now().Format("2006-01-02T00:00:00Z"),
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
	apiresp := c.getApiResponse(ctx)
	instance := models.TrainDetailsResponse{}
	json.Unmarshal(apiresp, &instance)

	train := instance.TrainSchedulerDetails[0]
	if tid := ctx.Query("train"); tid != "" {
		for _, detail := range instance.TrainSchedulerDetails {
			if detail.Train.TrainID == tid {
				fmt.Println(tid)
				train = detail
			}
		}
	}

	ctx.HTML(http.StatusOK, "traininfo/info", gin.H{
		"info":            train,
		"tid":             ctx.Query("tid"),
		"selectedTrainID": train.Train.TrainID,
		"trains":          instance.TrainSchedulerDetails,
		"numberOfTrains":  len(instance.TrainSchedulerDetails),
	})
}
