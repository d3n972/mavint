package models

import "encoding/json"

type Stations struct {
	Stations []struct {
		Id                            int    `json:"id"`
		IsAlias                       bool   `json:"isAlias"`
		Name                          string `json:"name"`
		Code                          string `json:"code"`
		BaseCode                      string `json:"baseCode"`
		IsInternational               bool   `json:"isInternational"`
		CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
		CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
		Country                       string `json:"country"`
		CoutryIso                     string `json:"coutryIso"`
		IsIn1081                      bool   `json:"isIn108_1"`
	} `json:"stations"`
}

func (s *Stations) Load(data []byte) error {
	return json.Unmarshal(data, s)
}
