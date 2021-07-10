package data

import (
	"encoding/json"
	"io"
)

type MessageRequest struct {

	Request   Request `json:"request" gorm:"embedded"`
	MessageID int     `json:"messageid" gorm:"uniqueIndex"`
	ForWho string `json:"forWho"`
}

func (p *MessageRequest) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// collection
type MessageRequests []*MessageRequest

// encode (using json new encoder over marshall)
func (p *MessageRequests) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetMessageRequests() MessageRequests {
	return messageRequestsList
}


var messageRequestsList = []*MessageRequest{

	{
		Request : Request{
			//ID: 2,
			SentBy : "wintzy",
			Status: DENIED,
		},
		MessageID: 123,
		ForWho: "nikola123",

	},
	{
		Request : Request{
			//ID: 2,
			SentBy : "dani",
			Status: DENIED,
		},
		MessageID: 123,
		ForWho: "tomik",
	},
}