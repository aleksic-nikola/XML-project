package data

import (
	"encoding/json"
	"io"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

// declaring the collection
type Users []*User

func GetUsers() Users {
	return userList
}

var userList = []*User{

	{
		ID: 1,
		Name: "Danilo",
		Lastname: "Paripovic",
		Username: "dparipovic",
		Email: "dparipovic@mail",
		Password: "aaaaa",
	},
}