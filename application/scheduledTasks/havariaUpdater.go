package scheduledTasks

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/d3n972/mavint/domain"
	"github.com/d3n972/mavint/domain/models"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func _apigetdata() models.MapTrainsResponse {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	type Jo struct {
		Pre     bool `json:"pre"`
		History bool `json:"history"`
		ID      bool `json:"id"`
	}
	type Payload struct {
		A  string `json:"a"`
		Jo Jo     `json:"jo"`
	}

	data := Payload{A: "TRAINS", Jo: Jo{Pre: true, History: false, ID: false}}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://vonatinfo.mav-start.hu/map.aspx/getData", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "en-US,en-GB;q=0.9,en;q=0.8,hu-HU;q=0.7,hu;q=0.6")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Cookie", os.ExpandEnv("f5avraaaaaaaaaaaaaaaa_session_=DNGKGCAOKJBEIEGJHKEDEDEBCLNGOECOIHAMFBOBBCNHHIDGMCAPOKKHMBPDEFHOFCLDENMLLEPIPPAJNLIANNGINDJKFHLCFIFHAIPFIFJIMNKDKIALCGBHCEDHBOGO; ASP.NET_SessionId=pvcwoawnt0k1ryl4b0g4mhy0; show_legend=true; dtCookie=v_4_srv_1_sn_09A4552A51A14F62B0AAA52DFF665497_perc_100000_ol_0_mul_1_app-3Aea7c4b59f27d43eb_1_rcs-3Acss_0; rxVisitor=1649745312848UIMN060IMV1MBD60GG4A3L7HCA73HL36; rxvt=1649747494986|1649745312850; dtLatC=11; dtSa=true%7CKU%7C-1%7CExample%7C-%7C1649745698608%7C345543504_841%7Chttps%3A%2F%2Fvim.mav-start.hu%2FVIM%2FPRMV%2F20211030%2FMobileService.svc%2Frest%2Fhelp%2Foperations%2FGetMozdonyInfo%7C%7C%7C%7C; dtPC=1$345543504_841h-vKANTNVRKIDHMHFOPDISUMURVHAKAHGNI-0e0; show_vim=1659537884771; f5avraaaaaaaaaaaaaaaa_session_=GLANHNODIJOOBLFCGJDOAPJDBIHHJLJAMNGKLDMGEPNELBPMNIBHJPBEFOIGFLNIDEEDPHHICKEFLCJNCGBAOHBBJKDPNFGLNEGGLODNPCEADDMLOIMJNNJKOMOGDLLD"))
	req.Header.Set("Origin", "http://vonatinfo.mav-start.hu")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "http://vonatinfo.mav-start.hu/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := client.Do(req)
	if err != nil {
		// handle err
	}
	respBytes, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	mtr := models.MapTrainsResponse{}
	json.Unmarshal(respBytes, &mtr)
	return mtr
}
func HavarianUpdaterTask() *domain.Schedule {
	return &domain.Schedule{
		Interval: 1 * time.Minute,
		Handler: func(ctx domain.AppContext) {
			R := ctx.Redis
			if R.Exists(context.TODO(), "havariaCache").Val() != 222 {
				hc := models.HavariaCache{}
				trains := _apigetdata().D.Result.Trains.Train
				for _, t := range trains {
					tn := strings.Replace(t.TrainNumber, "55", "", 1)
					hc[tn] = models.HavariaCacheEntry{Time: t.Delay}
				}
				b, e := json.Marshal(hc)
				if e != nil {
				}
				R.Set(context.TODO(), "havariaCache", b, 5*time.Minute)
			}
		},
	}
}
