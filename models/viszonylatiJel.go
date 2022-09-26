package models

type ViszonylatiJel struct {
	PiktogramFullName string `json:"piktogramFullName"`
	FontSzinKod       string `json:"fontSzinKod"`
	HatterSzinKod     string `json:"hatterSzinKod"`
	Sign              Sign   `json:"sign"`
	Jel               string `json:"jel"`
}
