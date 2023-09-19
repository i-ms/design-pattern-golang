package observer

import (
	"fmt"
)

type User struct {
	profileType string
	name        string
	followers   *ProfileList
}

func NewUser(name string) *User {
	return &User{
		profileType: "user",
		name:        name,
		followers:   NewProfileList(),
	}
}

func (u *User) CreatePost(m string) {
	fmt.Printf("%s created a new post : %s\n", u.name, m)
	u.Notify(m)
}

func (u *User) Notify(m string) {
	for _, f := range u.followers.profiles {
		f.Update(u.name, m)
	}
}

func (u *User) Update(name, message string) {
	fmt.Printf("User %s received an update from %s : %s\n", u.name, name, message)
}
