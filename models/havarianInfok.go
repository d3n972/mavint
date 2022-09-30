package models

type HavarianInfok struct {
	AktualisKeses float64  `json:"aktualisKeses"`
	KesesiOk      string   `json:"kesesiOk"`
	HavariaInfo   []string `json:"havariaInfo"`
	UzletiInfo    string   `json:"uzletiInfo"`
	KesesInfo     string   `json:"kesesInfo"`
}
