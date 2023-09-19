package observer

import (
	"fmt"
)

type Channel struct {
	profileType string
	name        string
	followers   *ProfileList
}

func NewChannel(name string) *Channel {
	return &Channel{
		profileType: "channel",
		name:        name,
		followers:   NewProfileList(),
	}
}

func (c *Channel) CreatePost(m string) {
	fmt.Printf("Channel %s created new post : %s\n", c.name, m)
	c.Notify(m)
}

func (c *Channel) Notify(m string) {
	for _, f := range c.followers.profiles {
		f.Update(c.name, m)
	}
}

func (c *Channel) Update(name, message string) {
	fmt.Printf("Channel %s received an update from %s : %s\n", c.name, name, message)
}
