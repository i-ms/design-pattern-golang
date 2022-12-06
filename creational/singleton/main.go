package singleton

import "fmt"

func App() {
	for i := 0; i < 10; i++ {
		go getInstance()
	}
	fmt.Scanln()
}
