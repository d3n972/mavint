package models

type Kind struct {
	Name                string       `json:"name"`
	SortName            string       `json:"sortName"`
	Code                string       `json:"code"`
	Priority            int          `json:"priority"`
	BackgroundColorCode string       `json:"backgroundColorCode"`
	ForegroundColorCode string       `json:"foregroundColorCode"`
	Sign                Sign         `json:"sign"`
	StartStation        StartStation `json:"startStation"`
	EndStation          EndStation   `json:"endStation"`
}
type Kinds struct {
	Kind
}
