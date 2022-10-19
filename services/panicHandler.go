package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var gt_dsn string

func PanicHandler(c *gin.Context, plusdata any) {
	if rec := recover(); rec != nil {
		if os.Getenv("GIN_MODE") != "release" {
			gt_dsn = "https://b622fdae7cdc49d3a036234ac3d0dfeb@gt.d3n.it/2"
		} else {
			gt_dsn = "https://c8d19ca3e2214dda92fd358c2d853029@gt.d3n.it/1"
		}
		sentry.Init(sentry.ClientOptions{
			Dsn: gt_dsn,
		})
		if err, ok := rec.(error); ok {
			sentry.CaptureException(err)
			sentry.Flush(2 * time.Second)
		} else {
			panic("error is not type of error")
		}
		// that recovery also handle XHR's
		// you need handle it
		if strings.ToLower(c.Request.Header.Get("X-Requested-With")) == "xmlhttprequest" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": rec,
			})
		} else {
			if os.Getenv("GIN_MODE") != "release" {
				fmt.Printf("[GIN-debug] %+v\n", rec)
				fmt.Printf("%s\n", debug.Stack())
				return
			}
			params := url.Values{}
			params.Add("api_dev_key", `a27677239273fe28e48fe0d4087a1fa0`)
			params.Add("api_user_key", "bcc663c029e394754e282b6cc5ff983c")
			params.Add("api_folder_key", "M5ZK65Ba")
			params.Add("api_paste_private", "2")
			params.Add("api_paste_expire_date", "N")
			params.Add("api_paste_name", fmt.Sprintf("%s %s", time.Now().Format(time.UnixDate), rec.(error).Error()))
			params.Add("api_paste_code", fmt.Sprintf("%s\n%s", rec.(error).Error(), string(debug.Stack())))
			params.Add("api_option", `paste`)
			body := strings.NewReader(params.Encode())

			req, err := http.NewRequest("POST", "https://pastebin.com/api/api_post.php", body)
			if err != nil {
				// handle err
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				// handle err
			}
			defer resp.Body.Close()
			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Errorf("%s", err)
			}
			traceURL := strings.Replace(string(respBody), "https://pastebin.com/", "https://pastebin.com/raw/", 1)
			// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

			// curl \
			//   -H "Content-Type: application/json" \
			//   -d '{"username": "test", "content": "hello"}' \
			//   $WEBHOOK_URL

			type Payload struct {
				Username string `json:"username"`
				Content  string `json:"content"`
			}

			data := Payload{
				Username: "Go Exception Handler",
				Content:  fmt.Sprintf("%s\n%s", rec.(error).Error(), traceURL),
			}
			payloadBytes, err := json.Marshal(data)
			if err != nil {
				// handle err
			}
			body2 := bytes.NewReader(payloadBytes)

			req, err = http.NewRequest("POST", "https://discord.com/api/webhooks/1021446205131214898/V0pKEg2xQykjRizJugmopAzQ6y8Y_pfTzgMpmA9rXicrAZw5_HDWtUXGYfxsZ7l6kzkW", body2)
			if err != nil {
				// handle err
			}
			req.Header.Set("Content-Type", "application/json")
			resp, err = http.DefaultClient.Do(req)
			if err != nil {
				// handle err
			}
			defer resp.Body.Close()
			c.HTML(http.StatusOK, "500", gin.H{
				"error":    rec.(error).Error(),
				"traceURL": traceURL,
			})
		}
	}
}
