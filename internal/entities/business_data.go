package entities

import (
	"strconv"
	"strings"
)

type BusinessData struct {
	SeriesReference string `json:"series_reference"`
	Period          string `json:"period"`
	DataValue       string `json:"data_value"`
	Suppressed      string `json:"suppressed"`
	Status          bool   `json:"status"`
	Units           string `json:"units"`
	Magnitude       string `json:"magnitude"`
	Subject         string `json:"subject"`
	Group           string `json:"group"`
	SeriesTitle1    string `json:"series_title1"`
	SeriesTitle2    string `json:"series_title2"`
	SeriesTitle3    string `json:"series_title3"`
	SeriesTitle4    string `json:"series_title4"`
	SeriesTitle5    string `json:"series_title5"`
}

func NewFromCSVLine(line string, delimiter string) (*BusinessData, error) {
	var err error
	array := strings.Split(line, delimiter)
	s := new(BusinessData)

	s.Period = array[0]
	s.SeriesReference = array[1]
	s.DataValue = array[2]
	s.Suppressed = array[3]
	status := array[4]
	s.Status, err = strconv.ParseBool(status)
	if err != nil {
		return nil, err
	}
	s.Units = array[5]
	s.Magnitude = array[6]
	s.Subject = array[7]
	s.Group = array[8]
	s.SeriesTitle1 = array[9]
	s.SeriesTitle2 = array[10]
	s.SeriesTitle3 = array[11]
	s.SeriesTitle4 = array[12]
	s.SeriesTitle5 = array[13]

	return s, nil
}
