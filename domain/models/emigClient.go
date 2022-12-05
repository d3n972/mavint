package models

import (
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type EmigClient struct {
	SqlId     string
	SessionId string
}

func NewEmigClient() *EmigClient {
	e := EmigClient{}
	e.SqlId = e.GetSQLID()
	e.SessionId = e.GetSessionID()
	return &e
}
func (e *EmigClient) GetSessionID() string {
	req, err := http.NewRequest("GET", "https://iemig.mav-trakcio.hu/emig7/emig.aspx?v=5.1", nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en-GB;q=0.9,en;q=0.8,hu-HU;q=0.7,hu;q=0.6,no;q=0.5")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://iemig.mav-trakcio.hu/7.0/indexsp.html?ModuleType=none&ModuleCode=none&ModuleKey=none&r=1661185507114")
	req.Header.Set("Sec-Fetch-Dest", "script")
	req.Header.Set("Sec-Fetch-Mode", "no-cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	req.Header.Set("Sec-Ch-Ua", "\".Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	d, _ := io.ReadAll(resp.Body)

	defer resp.Body.Close()
	x, _ := regexp.Compile("gSessionId=[0-9]{9}")
	o := string(x.FindAll(d, 1)[0])
	return strings.Split(o, "=")[1]
}
func (e *EmigClient) GetSQLID() string {
	sessID := e.GetSessionID()
	params := url.Values{}
	params.Add("u", `public`)
	params.Add("s", string(sessID))
	params.Add("t", `publicsandr`)
	params.Add("q", `Q5`)
	params.Add("lt", `SqlCreate`)
	params.Add("w", `null`)
	params.Add("c", `null`)
	params.Add("o", `null`)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://iemig.mav-trakcio.hu/emig7/emig.aspx", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en-GB;q=0.9,en;q=0.8,hu-HU;q=0.7,hu;q=0.6,no;q=0.5")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "https://iemig.mav-trakcio.hu")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://iemig.mav-trakcio.hu/7.0/indexsp.html?ModuleType=none&ModuleCode=none&ModuleKey=none&r=1661185507114")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Sec-Ch-Ua", "\".Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	data, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	x, _ := regexp.Compile("<sqlid>[0-9]*")
	sqlid := string(x.FindAll(data, 1)[0])
	return strings.Replace(sqlid, "<sqlid>", "", 1)
}
