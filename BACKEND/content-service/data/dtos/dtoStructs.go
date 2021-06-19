package dtos

import (
	"encoding/json"
	"io")


type UsernameRole struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (ur *UsernameRole) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ur)
}

func (ur *UsernameRole) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ur)
}
