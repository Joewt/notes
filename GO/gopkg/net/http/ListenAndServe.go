package main

import (
	"io"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello")
}

func main() {
	http.HandleFunc("/hello", Hello)
	err := http.ListenAndServeTLS(":8080", "joe.crt", "joe.key", nil)
	if err != nil {
		panic(err)
	}
}
