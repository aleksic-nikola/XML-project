package data

import (
	"encoding/json"
	"io"
)

type FollowRequest struct {

	RequestID uint `json:"request_id"`
	Request Request `json:"request" gorm:"foreignkey=RequestID"`
	For string `json:"for"`

}


func (p *FollowRequest) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// collection
type FollowRequests []*FollowRequest

// encode (using json new encoder over marshall)
func (p *FollowRequests) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}



func GetFollowRequests() FollowRequests {
	return followRequestList
}


var followRequestList = []*FollowRequest{

	{
		RequestID: 2,
		Request : Request{
			//ID: 2,
			SentBy : "wintzy",
			Status: ACCEPTED,
		},
		For: "nikola123",

	},
	{
		RequestID: 3,
		Request : Request{
			//ID: 3,
			SentBy : "dani",
			Status: INPROCESS,
		},
		For: "tomik333",
	},
}