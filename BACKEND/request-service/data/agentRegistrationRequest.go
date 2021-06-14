package data

import (
	"encoding/json"
	"io"
)

type AgentRegistrationRequest struct {

	RequestID uint `json:"request_id"`
	Request Request `json:"request" gorm:"foreignkey=RequestID"`
	Name string `json:"name"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
	Email string `json:"email"`
	Website string `json:"website"`

}

func (p *AgentRegistrationRequest) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// collection
type AgentRegistrationRequests []*AgentRegistrationRequest

// encode (using json new encoder over marshall)
func (p *AgentRegistrationRequests) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetAgentRegistrationRequests() AgentRegistrationRequests {
	return agentRegistrationRequestList
}


var agentRegistrationRequestList = []*AgentRegistrationRequest{

	{
		RequestID: 1,
		Request : Request{
			//ID: 2,
			SentBy : "wintzy",
			Status: DENIED,
		},
		Name: "Pera",
		LastName: "Peric",
		Password: "12345",
		Email: "pera@gmail.com",
		Website: "www.pera.com",

	},
	{
		RequestID: 2,
		Request : Request{
			//ID: 2,
			SentBy : "dani",
			Status: DENIED,
		},
		Name: "Jovan",
		LastName: "Jovanovic",
		Password: "12345",
		Email: "jovan@gmail.com",
		Website: "www.jovan.com",
	},
}