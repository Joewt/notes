package main

import (
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	//http.Error(w, "http error", 404)
	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/hello", Hello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
