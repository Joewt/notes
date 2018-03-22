package main

import "fmt"

func main() {
	s := "\u0099\u7455"
	for i, c := range s {
		fmt.Printf("%d:%c", i, c)
	}
}