package option

import (
	"fmt"
	"testing"
)

func TestUserOption(t *testing.T) {
	u := NewUser(
		WithId(1),
		WithName("John"),
		WithEmail("sample@example.com"),
		WithDateOfBirth("1990-01-01"),
		WithPhone("1234567890"))

	fmt.Println(u)
}
