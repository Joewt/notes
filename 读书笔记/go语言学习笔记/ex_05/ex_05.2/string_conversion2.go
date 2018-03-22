package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "10"
	var newS string
	fmt.Printf("%d\n", strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not a integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The inerger is %d\n", an)
	an += 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}