package data

import (
	"encoding/json"
	"io"
)

type SensitiveContentReportRequest struct {

	Request Request `json:"request"`
	PostID string `json:"postID"`
	Note string `json:"note"`

}

func (p *SensitiveContentReportRequest) FromJSON(r io.Reader) error {
e := json.NewDecoder(r)
return e.Decode(p)
}

// collection
type SensitiveContentReportRequests []*SensitiveContentReportRequest

// encode (using json new encoder over marshall)
func (p *SensitiveContentReportRequests) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetSensitiveContentReportRequests() SensitiveContentReportRequests {
	return SensitiveContentReportRequestList
}

var SensitiveContentReportRequestList = []*SensitiveContentReportRequest{

	{
		Request : Request{
			ID: 2,
			SentBy : "wintzy",
			Status: DENIED,
		},
		PostID: "34356",
		Note: "This is explicit content because...",

	},
	{
		Request : Request{
			ID: 4,
			SentBy : "pera123",
			Status: ACCEPTED,
		},
		PostID: "11256",
		Note: "This is explicit content because.....",

	},
}