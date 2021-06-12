package data

import (
	"encoding/json"
	"io"
)

type MonitorReportRequest struct {

	RequestID uint `json:"request_id"`
	Request   Request `json:"request" gorm:"foreignkey=RequestID"`
	ForCampaign int `json:"forCampaign"`

}

func (p *MonitorReportRequest) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// collection
type MonitorReportRequests []*MonitorReportRequest

// encode (using json new encoder over marshall)
func (p *MonitorReportRequests) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetMonitorReportRequests() MonitorReportRequests {
	return monitorReportRequestList
}


var monitorReportRequestList = []*MonitorReportRequest{

	{
		RequestID: 3,
		Request : Request{
			//ID: 2,
			SentBy : "wintzy",
			Status: DENIED,
		},
		ForCampaign: 123,

	},
	{
		RequestID: 3,
		Request : Request{
			//ID: 2,
			SentBy : "dani",
			Status: ACCEPTED,
		},
		ForCampaign: 1213,
	},
}
