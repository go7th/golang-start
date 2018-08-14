// test project main.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	var (
		aa string = "Hello World122!"
	)
	http.HandleFunc("/", func(
		writer http.ResponseWriter,
		request *http.Request) {
		fmt.Fprintf(writer,
			"<h1>hello world %s!</h1>", aa)
	})
	http.ListenAndServe(":80", nil)

}
