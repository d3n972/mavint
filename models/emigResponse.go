package models

import (
	"encoding/xml"
	"strconv"
)

type Mozdony struct {
	Text      string `xml:",chardata"`
	ID        string `xml:"id,attr"`
	Lat       string `xml:"lat,attr"`
	Lng       string `xml:"lng,attr"`
	Icon      string `xml:"icon,attr"`
	Title     string `xml:"title,attr"`
	Tipus     string `xml:"tipus,attr"`
	Vonatszam string `xml:"vonatszam,attr"`
	Uic       string `xml:"uic,attr"`
}

func (m *Mozdony) GetCoords() []float64 {
	lan, _ := strconv.Atoi(m.Lat)
	lon, _ := strconv.Atoi(m.Lng)
	return []float64{
		float64(lan) / 1000000,
		float64(lon) / 1000000,
	}
}

type EmigResponse struct {
	XMLName   xml.Name `xml:"response"`
	Text      string   `xml:",chardata"`
	Mozdonyok struct {
		Text    string    `xml:",chardata"`
		ID      string    `xml:"id,attr"`
		Mozdony []Mozdony `xml:"Mozdony"`
	} `xml:"mozdonyok"`
	Asqf      string `xml:"asqf"`
	Copyright string `xml:"copyright"`
	Transfer  struct {
		Text  string `xml:",chardata"`
		Delay string `xml:"delay,attr"`
		Color string `xml:"color,attr"`
	} `xml:"transfer"`
	Status struct {
		Text      string `xml:",chardata"`
		Code      string `xml:"code,attr"`
		Mozdonyok string `xml:"mozdonyok,attr"`
	} `xml:"status"`
}
