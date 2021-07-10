package data

import (
	"encoding/json"
	"gorm.io/gorm"
	"io"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username" gorm:"uniqueIndex"`
	Email    string `json:"email" gorm:"uniqueIndex"`
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
		Name: "Danilo",
		Lastname: "Paripovic",
		Username: "dparipovic",
		Email: "dparipovic@mail",
		Password: "aaaaa",
	},
}