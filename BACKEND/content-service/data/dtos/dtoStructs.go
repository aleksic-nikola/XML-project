package dtos

import (
	"encoding/json"
	"io"
)


type UsernameRole struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UsernameDto struct {
	Username string `json:"username"`
}

type UserIDDto struct{
	UserId int `json:"user_id"`
}

func (ur *UserIDDto) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ur)
}

func (ur *UsernameRole) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ur)
}

func (ur *UsernameRole) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ur)
}

func (ur *UsernameDto) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ur)
}

