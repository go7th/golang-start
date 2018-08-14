// test project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{3, 9, 4, 5, 7, 46, 23}
	sort.Ints(a)
	for _, v := range a {
		fmt.Println(v)
		//		fmt.Println(i)
	}
}
