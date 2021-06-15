package dto

import (
	"encoding/json"
	"io"
)

type UserEditDTO struct {
	OldUsername string `json:"oldusername"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}

func (u *UserEditDTO) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func (u *UserEditDTO) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}