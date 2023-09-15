package iteratorsocial

import "testing"

func TestSocialIterator(t *testing.T) {
	facebook := &Facebook{}
	spammer := &SocialSpammer{}

	// Create iterator for friends and coworkers.
	friendIterator := facebook.CreateFriendsIterator(1)
	coworkerIterator := facebook.CreateCoworkersIterator(1)

	// send spam message to friends
	println("Sending spam message to friends")
	spammer.SendSpam(friendIterator, "Hello friends!")

	// send spam message to coworkers
	println("Sending spam message to coworkers")
	spammer.SendSpam(coworkerIterator, "Hello coworkers!")
}
