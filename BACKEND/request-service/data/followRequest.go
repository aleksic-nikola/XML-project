package data

import (
	"encoding/json"
	"io"
)

type FollowRequest struct {

	Request Request `json:"request" gorm:"embedded"`
	ForWho string `json:"forWho"`

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
		Request : Request{
			//ID: 2,
			SentBy : "wintzy",
			Status: ACCEPTED,
		},
		ForWho: "nikola123",

	},
	{

		Request : Request{
			//ID: 3,
			SentBy : "dani",
			Status: INPROCESS,
		},
		ForWho: "tomik333",
	},
}