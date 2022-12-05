package scheduledTasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/d3n972/mavint/domain"
	"github.com/d3n972/mavint/domain/dao"
	"github.com/d3n972/mavint/domain/models"
	"github.com/d3n972/mavint/domain/repository"
	"io"
	"net/http"
	"strings"
	"time"
)

type VPELoggerTask struct {
	interval time.Duration
}

func (v VPELoggerTask) getData() models.VPEResponse {
	body := strings.NewReader("default_orderby=vpe_azonosito&default_orderby_asc=false&pager_limit=5000&pager_offset=0&szuro_id=&with_count=false")
	req, err := http.NewRequest("POST", "https://www.kapella2.hu/vaganyzar/vaganyzarlistazo/lista", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "en-US,en-GB;q=0.9,en;q=0.8,hu-HU;q=0.7,hu;q=0.6")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", "PROD=vs8f7sdfsdft3434p8vch15ntqac4; k2_language_cookie=hu")
	req.Header.Set("Origin", "https://www.kapella2.hu")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://www.kapella2.hu/vaganyzar/vaganyzarlistazo")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"106\", \"Google Chrome\";v=\"106\", \"Not;A=Brand\";v=\"99\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	fmt.Printf("VPE: %s", resp.Status)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	e := models.VPEResponse{}
	json.Unmarshal(data, &e)
	return e
}
func (v VPELoggerTask) Handler(ctx domain.AppContext) {
	time.LoadLocation("Europe/Budapest")
	ctx.Db.AutoMigrate(&models.VPEEntry{})
	vpedata := v.getData()
	for _, entry := range vpedata.Data {
		fromTime, _ := time.Parse("2006-01-02 15:04:05", entry.ErvenyessegKezdoWots)
		untilTime, _ := time.Parse("2006-01-02 15:04:05", entry.ErvenyessegVegWots)
		record := models.VPEEntry{
			VPEHash:   entry.Hsh,
			From:      fromTime,
			Until:     untilTime,
			Provider:  entry.SzNev,
			VPEID:     entry.VpeAzonosito,
			EntryType: entry.IgenylesTipusNev,
			VPEName:   entry.VuNevTooltip,
			Status:    entry.StatuszKod,
		}
		c := int64(0)
		repo := repository.NewRepository[dao.VPEEntryDAO, models.VPEEntry](context.Background())
		res, er := repo.Find(context.Background(),
			repository.NewEqualsSpecification(
				"vpe_hash", entry.Hsh,
			),
		)
		if er != nil {
			//todo: handle error case
		}
		c = int64(len(res))
		if c == 0 {
			fmt.Printf("[VPE] Saving: %s\n", entry.VpeAzonosito)
			repo.Insert(context.Background(), &record)
		}
	}
}
