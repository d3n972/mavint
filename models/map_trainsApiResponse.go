package models

type MapTrainsResponse struct {
	D D `json:"d"`
}
type Param struct {
	Pre     bool `json:"pre"`
	History bool `json:"history"`
	ID      bool `json:"id"`
}
type Train struct {
	Delay       int     `json:"@Delay"`
	Lat         float64 `json:"@Lat"`
	Relation    string  `json:"@Relation"`
	TrainNumber string  `json:"@TrainNumber"`
	Menetvonal  string  `json:"@Menetvonal"`
	Line        string  `json:"@Line,omitempty"`
	Lon         float64 `json:"@Lon"`
	ElviraID    string  `json:"@ElviraID"`
}
type Trains struct {
	Train []Train `json:"Train"`
}
type Result struct {
	PackageType  string `json:"@PackageType"`
	CreationTime string `json:"@CreationTime"`
	Trains       Trains `json:"Trains"`
}
type D struct {
	Type   string `json:"__type"`
	Action string `json:"action"`
	Param  Param  `json:"param"`
	Result Result `json:"result"`
	Time   string `json:"time"`
}
