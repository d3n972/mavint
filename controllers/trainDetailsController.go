package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/d3n972/mavint/scheduledTasks"
	"github.com/d3n972/mavint/services"
	"github.com/go-redis/redis/v9"
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
	var cacheId string
	var respBytes []byte
	if ctx.Params.ByName("train") != "" {
		i, _ := strconv.Atoi(ctx.Params.ByName("train"))
		cacheId = "train:" + strconv.Itoa(i)
		data = Payload{
			MaxCount:   "9999999",
			MinCount:   "0",
			TrainID:    i,
			TravelDate: time.Now().Format("2006-01-02T00:00:00Z"),
			Type:       "TrainInfo",
		}
	} else {
		cacheId = "train:" + ctx.Query("tid")
		data = Payload{
			MaxCount:    "9999999",
			MinCount:    "0",
			TrainNumber: ctx.Query("tid"),
			TravelDate:  time.Now().Format("2006-01-02T00:00:00Z"),
			Type:        "TrainInfo",
		}
	}
	hRdb, _ := ctx.Get("cache")
	rdb := hRdb.(*redis.Client)
	if cnt, _ := rdb.Exists(context.Background(), cacheId).Result(); cnt == 0 {

		payloadBytes, err := json.Marshal(data)
		fmt.Printf("payloadBytes: %v\n", string(payloadBytes))
		if err != nil {
			panic(err)
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
		req.Header.Set("Usersessionid", "4d793891-0e70-45b5-a9c6-64eddbba2532")
		req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"104\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"104\"")
		req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
		req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		res, _ := io.ReadAll(resp.Body)
		fmt.Printf("Saving response for key %s\n", cacheId)
		rdb.Set(context.Background(), cacheId, res, 5*time.Minute)
		respBytes, _ = rdb.Get(context.Background(), cacheId).Bytes()
	} else {
		fmt.Printf("Using cached response for key %s\n", cacheId)
		respBytes, _ = rdb.Get(context.Background(), cacheId).Bytes()
	}
	return respBytes
}
func (c *TrainDetailsController) Render(ctx *gin.Context) {
	apiresp := c.getApiResponse(ctx)
	check := gin.H{}
	instance := models.TrainDetailsResponse{}
	json.Unmarshal(apiresp, &instance)
	json.Unmarshal(apiresp, &check)
	if instance.TrainSchedulerDetails == nil {
		//we have an error!
		ctx.HTML(http.StatusNotFound, "pages/article", gin.H{
			"title":   instance.ExceptionMessage,
			"pub":     "",
			"content": []interface{}{},
		})
	}
	train := instance.TrainSchedulerDetails[0]
	if tid := ctx.Query("train"); tid != "" {
		for _, detail := range instance.TrainSchedulerDetails {
			if detail.Train.TrainID == tid {
				train = detail
			}
		}
	}

	ed := services.EngineDiscovery{}
	appctx, _ := ctx.Get("appctx")
	engine, _ := ed.FindByTrainNumber(context.WithValue(context.Background(), "db", appctx.(scheduledTasks.AppContext).Db), train.Train.Code)

	ctx.HTML(http.StatusOK, "traininfo/info_next", gin.H{
		"info":            train,
		"tid":             ctx.Query("tid"),
		"selectedTrainID": train.Train.TrainID,
		"trid":            ctx.Query("train"),
		"engineUIC":       engine.UIC,
		"currDate":        time.Now().Format("2006-01-02"),
		"trains":          instance.TrainSchedulerDetails,
		"numberOfTrains":  len(instance.TrainSchedulerDetails),
	})
}
func (c *TrainDetailsController) ESDDisplay(ctx *gin.Context) {
	apiresp := c.getApiResponse(ctx)
	instance := models.TrainDetailsResponse{}
	json.Unmarshal(apiresp, &instance)
	staytime := 0 * time.Minute
	var train models.TrainSchedulerDetails
	train = instance.TrainSchedulerDetails[0]
	if tid := ctx.Query("train"); tid != "" {
		for _, detail := range instance.TrainSchedulerDetails {
			if detail.Train.TrainID == tid {
				train = detail
			}
		}
	}
	for _, stop := range train.Scheduler {
		if stop.Start != nil && stop.Arrive != nil {
			delta := stop.Start.Sub(*stop.Arrive)
			staytime = staytime + delta
		}
	}
	stayT := time.Time{}
	stayT = stayT.Add(staytime)
	ctx.HTML(http.StatusOK, "pages/esd_display", gin.H{
		"schedule": train.Scheduler,
		"train":    train.Train,
		"trainID":  ctx.Query("tid"),
		"genDate":  time.Now().Local().Format("06.01.02."),
		"genTime":  time.Now().Local().Format("15:04"),
		"staytime": stayT,
	})
}
