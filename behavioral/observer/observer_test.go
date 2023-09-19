package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {
	alice := NewUser("Alice")
	bob := NewUser("Bob")
	charlie := NewUser("Charlie")

	fitness := NewChannel("Fitness")
	education := NewChannel("Education")
	worldAffairs := NewChannel("World Affairs")

	alice.followers.Add(fitness)
	alice.followers.Add(bob)

	education.followers.Add(alice)

	worldAffairs.followers.Add(charlie)

	alice.CreatePost("My first post")

	education.CreatePost("My first education post")

	worldAffairs.CreatePost("My first world affairs post")
}
