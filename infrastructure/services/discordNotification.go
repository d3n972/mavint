package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DiscordNotification struct {
	INotification
	whId       string
	whToken    string
	whNickname string
}

func (d *DiscordNotification) Init(p *Params) (bool, error) {
	d.whId = (*p)["id"]
	d.whToken = (*p)["token"]
	d.whNickname = (*p)["name"]
	return true, nil
}
func (d *DiscordNotification) Send(v any) error {
	type Payload struct {
		Username string `json:"username"`
		Content  string `json:"content"`
	}

	data := Payload{
		Username: d.whNickname,
		Content:  fmt.Sprintf("%s", v),
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", fmt.Sprintf("https://discord.com/api/webhooks/%s/%s", d.whId, d.whToken), body)
	print(fmt.Sprintf("https://discord.com/api/webhooks/%s/%s", d.whId, d.whToken))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
