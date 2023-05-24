package singleton

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	for i := 0; i < 10; i++ {
		go getInstance()
	}
	_, _ = fmt.Scanln()
}
