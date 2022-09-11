package controllers

import (
	"encoding/xml"
	"fmt"
	"github.com/d3n972/mavint/models"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type EmigController struct{}

func (e *EmigController) getData() models.EmigResponse {
	creds := models.NewEmigClient()
	fmt.Printf("%+v\n", creds)
	params := url.Values{}
	params.Add("u", `public`)
	params.Add("s", creds.SessionId)
	params.Add("t", `publicrspec`)
	params.Add("q", creds.SqlId)
	params.Add("f", `publicmlist`)
	//fmt.Printf("%s", params.Encode())
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://iemig.mav-trakcio.hu/emig7/emig.aspx", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en-GB;q=0.9,en;q=0.8,hu-HU;q=0.7,hu;q=0.6")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "https://iemig.mav-trakcio.hu")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://iemig.mav-trakcio.hu/7.0/indexsp.html?ModuleType=none&ModuleCode=none&ModuleKey=none&r=1662909593049")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"104\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"104\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")

	resp, err := http.DefaultClient.Do(req)
	respbytes, _ := io.ReadAll(resp.Body)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	er := models.EmigResponse{}
	xml.Unmarshal(respbytes, &er)
	fmt.Printf("%+v", er)
	return er
}
func (e *EmigController) GetTrainEngines(ctx *gin.Context) {
	emigResp := e.getData()
	ctx.JSON(http.StatusOK, gin.H{
		"e": emigResp.Mozdonyok,
	})
}
func (e *EmigController) Render(ctx *gin.Context) {
	emigResp := e.getData()
	ctx.JSON(http.StatusOK, gin.H{
		"e": emigResp.Mozdonyok,
	})
}
