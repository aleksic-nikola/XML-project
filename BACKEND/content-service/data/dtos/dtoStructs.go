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

type LikeDto struct {
	PostID int `json:"post_id"`
}

type CommentDto struct {
	Text string `json:"text"`
}

func (c *CommentDto) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
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

func (ur *UsernameDto) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ur)
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

