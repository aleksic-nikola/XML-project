package data

import (
	"encoding/json"
	"gorm.io/gorm"

	//"fmt"
	"io"
	//"time"
)

type Request struct {
	gorm.Model
	SentBy string `json:"sentby" gorm:"primaryKey"`
	Status RequestStatus `json:"requeststatus"`
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

var requestsList = []*Request{}
