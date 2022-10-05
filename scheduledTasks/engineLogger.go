package scheduledTasks

import (
	"encoding/xml"
	"github.com/d3n972/mavint/models"
	"github.com/d3n972/mavint/models/db"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func emig_getData() models.EmigResponse {
	creds := models.NewEmigClient()
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
	return er
}
func EngineLoggerTask() *Schedule {
	return &Schedule{
		Interval: 1 * time.Minute,
		Handler: func(ctx AppContext) {
			time.LoadLocation("Europe/Budapest")
			ctx.Db.AutoMigrate(db.EngineWorkday{})
			emigData := emig_getData()
			for _, engine := range emigData.Mozdonyok.Mozdony {
				tx := ctx.Db.Where(
					"ui_c = ? AND job_type = ? AND train_number = ? AND date = ?",
					engine.Uic,
					engine.Tipus,
					engine.Vonatszam,
					time.Now().Format("2006-01-02"),
				).Find(&db.EngineWorkday{})
				if tx.RowsAffected == 0 /*we don't have yet*/ {
					logEntry := db.EngineWorkday{

						UIC:         engine.Uic,
						Date:        time.Now().Format("2006-01-02"),
						JobType:     engine.Tipus,
						TrainNumber: &engine.Vonatszam,
					}
					logEntry.CreatedAt = time.Now().Local()
					logEntry.UpdatedAt = time.Now().Local()
					ctx.Db.Save(&logEntry)
				}
			}
		},
	}
}