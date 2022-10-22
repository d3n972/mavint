package models

type StopService struct {
	ListOrder                   interface{} `json:"listOrder"`
	Description                 interface{} `json:"description"`
	RestrictiveStartStationCode interface{} `json:"restrictiveStartStationCode"`
	RestrictiveEndStationCode   interface{} `json:"restrictiveEndStationCode"`
	Sign                        interface{} `json:"sign"`
	TrainStopKind               interface{} `json:"trainStopKind"`
}
