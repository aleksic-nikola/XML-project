package data

import (
	"encoding/json"
	//"fmt"
	"io"
	//"time"
)

type Request struct {

	ID int `json:"id"`
	SentBy string `json:"sentby"`
	Status RequestStatus `json:"RequestStatus"`
}

func (p *Request) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// collection
type Requests []*Request

// encode (using json new encoder over marshall)
func (p *Requests) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetRequests() Requests {
	return requestsList
}

var requestsList = []*Request{

	{
		ID: 1,
		SentBy: "lucyxz",
		Status: ACCEPTED,

	},
	{
		ID: 2,
		SentBy : "wintzy",
		Status: DENIED,
	},
}
