package dto

import (
	"encoding/json"
	"io"
)

type ProfileEditDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Phone string `json:"phone"`
	Gender int `json:"gender"`
	DateOfBirth string `json:"dateofbirth"`
	Website string `json:"website"`
	Biography string `json:"biography"`
}

type UserEditDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}

func (p *ProfileEditDTO) ProfFromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *ProfileEditDTO) ProfToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (u *UserEditDTO) UserFromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func (u *UserEditDTO) UserToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

type UsernameRole struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (ur *UsernameRole) URFromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ur)
}

func (ur *UsernameRole) URToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ur)
}

type ProfilePublic struct {
	Public bool `json:"ispublic"`
}

func (p *ProfilePublic) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *ProfilePublic) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}