package data

import (
	"encoding/json"
	"io"

	"gorm.io/gorm"
)


type User struct {
	// id, email, username, name, lastname, password, role
	//ID int `json:"id"`
	gorm.Model
	Email string `json:"email" gorm:"uniqueIndex"`
	Username string `json:"username" gorm:"uniqueIndex"`
	Name string `json:"name"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
	Role string `json:"role"`	
}

type LoginForm struct {
	Username string
	Password string
}

func (lf *LoginForm) LFFromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(lf)
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
		Email: "lucyxz@gmail.com",
		Username: "lucyxz",
		Name: "Mark",
		LastName: "Ristic",
		Password: "1337",
		Role : "user",
	},
}


