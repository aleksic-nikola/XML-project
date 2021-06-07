package data

import (
	"encoding/json"
	"io")


type User struct {
	// id, email, username, name, lastname, password, role
	ID int `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
	Name string `json:"name"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
	Role string `json:"role"`
	
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
		Email: "lucyxz@gmail.com",
		Username: "lucyxz",
		Name: "Mark",
		LastName: "Ristic",
		Password: "1337",
		Role : "user",
	},
}

