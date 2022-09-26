package models

type Services struct {
	ListOrder                   int    `json:"listOrder"`
	Description                 string `json:"description"`
	RestrictiveStartStationCode string `json:"restrictiveStartStationCode"`
	RestrictiveEndStationCode   string `json:"restrictiveEndStationCode"`
	Sign                        Sign   `json:"sign"`
	TrainStopKind               string `json:"trainStopKind"`
}
