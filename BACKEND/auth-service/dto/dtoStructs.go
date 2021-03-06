package dto

import (
	"encoding/json"
	"io")


type UsernameRoleDto struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UserIdDto struct {
	UserId uint `json:"user_id"`
}

func (u *UserIdDto) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}


func (u *UsernameRoleDto) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

type TokenDto struct {
	Token string `json:"token"`
}

func (t *TokenDto) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(t)
}

func (u *UsernameRoleDto) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}