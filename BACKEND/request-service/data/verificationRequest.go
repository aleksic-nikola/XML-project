package data

import (
	"encoding/json"
	"io"
)

type VerificationRequest struct {

	Request   Request `json:"request" gorm:"embedded"`
	Category VerifiedType `json:"verifiedType"`
	Image string `json:"image"`
	Name string `json:"name"`
	LastName string `json:"lastname"`

}


func (p *VerificationRequest) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// collection
type VerificationRequests []*VerificationRequest

// encode (using json new encoder over marshall)
func (p *VerificationRequests) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetVerificationRequests() VerificationRequests {
	return verificationRequestList
}


var  verificationRequestList = []*VerificationRequest{}
/*
	{
		Request : Request{
			//ID: 2,
			SentBy : "dani321",
			Status: INPROCESS,
		},
		Category: BRAND,
		Image: "../../img1",
		Name : "Nikola",
		LastName: "Nikolic",


	},
	{
		RequestID: 2,
		Request : Request{
			//ID: 2,
			SentBy : "pero321",
			Status: ACCEPTED,
		},
		Category: BUISNESS,
		Image: "../../img3",
		Name : "Marija",
		LastName: "Petrovic",
	},
}
*/