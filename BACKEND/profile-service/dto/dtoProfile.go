package dto

import (
	"encoding/json"
	"io"
	"xml/profile-service/data"
)


type UsernameRoleDto struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UsernameDto struct{
	Username string `json:"username"`
}

func (ur *UsernameDto) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ur)
}

type ProfileForFollow struct {
	FollowToUsername string `json:"followToUsername"`
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

type NewVerified struct {
	Username string `json:"username"`
	VerifiedType data.VerifiedType `json:"verified_type"`
}

func (nv *NewVerified) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(nv)
}

func (nv *NewVerified) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(nv)
}

type PostToFavourites struct {
	CollectionName string `json:"collection_name"`
	PostId         uint   `json:"post_id"`
}

func (ptf *PostToFavourites) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ptf)
}

func (ptf *PostToFavourites) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ptf)
}

type PostIdsDto struct {
	Ids []PostIdDto `json:"ids"`
}

type PostIdDto struct {
	Id uint `json:"id"`
}

func (ptf *PostIdsDto) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ptf)
}

func (ptf *PostIdsDto) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ptf)
}

func (ptf *PostIdDto) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ptf)
}

func (ptf *PostIdDto) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ptf)
}
