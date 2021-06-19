package dto

import (
	"encoding/json"
	"io")


type UsernameRoleDto struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UsernameFollowerDto struct{
	Username string `json:"username"`

}

type ProfileForFollow struct {
	FollowToUsername string `json:"follow-to-username"`
}

type RequestDto struct{
	SentBy string `json:"sentby"`
}

type FollowRequestDto struct{
	Request RequestDto `json:"request"`
	ForWho string `json:"forWho"`
}

func (u *FollowRequestDto) ToJSON(w io.Writer) error {
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

func (ur *UsernameRoleDto) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ur)
}