package iteratorsocial

import "fmt"

type SocialSpammer struct{}

func (ss *SocialSpammer) SendSpam(iterator ProfileIterator, message string) {
	for iterator.HasNext() {
		profile := iterator.Next()
		fmt.Printf("Sent message %s to email %s\n", message, profile.Email)
	}
}
