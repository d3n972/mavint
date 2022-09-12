package controllers

import (
	"encoding/xml"
	"github.com/PuerkitoBio/goquery"
	"github.com/d3n972/mavint/models"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strings"
)

type NewsController struct{}

func (n *NewsController) apiGetData() *models.RssFeed {
	req, err := http.NewRequest("GET", "https://www.mavcsoport.hu/mavinform/rss.xml", nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/jxl,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "en-US,en-GB;q=0.9,en;q=0.8,hu-HU;q=0.7,hu;q=0.6")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "f5avraaaaaaaaaaaaaaaa_session_=JAGHLAPLCBNFIAECKNHFLCLCDKDPBNLFAGDHPCINDHOMBHOELALOEDNCLIMJELFGFICDALGEHMEDGAEEALJANLJKAHBMPFKJBDHNPFEAIJDLMFBJLJIIJIJOICDKLLFJ; cookie-agreed-version=1.0.0; f5avraaaaaaaaaaaaaaaa_session_=KPNOGMLBOEMOFJNGFGLMGDNEODCDJJFOECDPOKCBHBFNFEPKAGOJMGOADAIDHGHOCHBDJDNODMEFPDFLLGHAMKFCAHHACMEDFJAIPILEDPGOEAPFBCHBICBAHAFCBMFD")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://www.mavcsoport.hu/")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"104\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"104\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	respBytes, _ := io.ReadAll(resp.Body)
	rf := models.RssFeed{}
	xml.Unmarshal(respBytes, &rf)
	return &rf
}
func (n *NewsController) Render(ctx *gin.Context) {
	rss := n.apiGetData()
	ctx.HTML(http.StatusOK, "pages/news", gin.H{
		"news": rss,
	})
}
func (n *NewsController) RenderArticle(ctx *gin.Context) {
	webPage := "https://www.mavcsoport.hu/print/" + ctx.Query("id") //ex: 109917
	resp, err := http.Get(webPage)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find("title").Text()
	publication := doc.Find(".node-date").Text()
	content := doc.Find(".field-body").Text()
	ctx.HTML(http.StatusOK, "pages/article", gin.H{
		"title":   title,
		"pub":     publication,
		"content": strings.Split(content, "\n"),
	})
}
