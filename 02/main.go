// 02 project main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 500; i++ {
		go OutInfo(i, ch)
	}
	for i := 0; i < 200; i++ {
		msg := <-ch
		fmt.Print(msg)
	}
	time.Sleep(5 * time.Millisecond)
}
func OutInfo(i int, ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("hello world %d! \n", i)
	}
}
