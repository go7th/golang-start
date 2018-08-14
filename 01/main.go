// test project main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5000; i++ {
		go OutInfo(i)
	}
	time.Sleep(5 * time.Millisecond)
}
func OutInfo(i int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("hello world %d! \n", i)
	}
}
