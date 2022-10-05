package models

type StationsResponse struct {
	Id                            int    `json:"id"`
	IsAlias                       bool   `json:"isAlias"`
	Name                          string `json:"name"`
	Code                          string `json:"code"`
	BaseCode                      string `json:"baseCode"`
	IsInternational               bool   `json:"isInternational"`
	CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
	CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
	Country                       string `json:"country"`
	CountryIso                    string `json:"coutryIso"`
	IsIn1081                      bool   `json:"isIn108_1"`
}
