package data

import (
	"encoding/json"
	"io"
	"time"
)

type MonitoringReport struct {
	ID int `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Campaign int `json:"campaign"`
	Likes int `json:"likes"`
	Dislikes int `json:"dislikes"`
	Comments int `json:"comments"`
	Placements int `json:"placements"`
	SentBy map[string]int `json:"sentby"`

}

func (mr *MonitoringReport) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(mr)
}

func (mr *MonitoringReports) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(mr)
}

type MonitoringReports []*MonitoringReport

func GetMonitoringReports() MonitoringReports {
	return monitoringReportList
}

var monitoringReportList = []*MonitoringReport {

	{
		ID: 1,
		Timestamp: time.Now(),
		Campaign: 1337,
		Likes: 37,
		Dislikes: 37,
		Comments: 37,
		Placements: 50,
		SentBy: map[string]int{
			"rsc": 3711,
			"r":   2138,
			"gri": 1908,
			"adg": 912,
		},
	},
}