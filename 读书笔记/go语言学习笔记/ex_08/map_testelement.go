package main

import "fmt"

func main() {
	var value int

	var isPresent bool

	map1 := make(map[string]int)
	map1["New Delhi"] = 25
	map1["Beijing"] = 20
	map1["Washington"] = 25
	value, isPresent = map1["Beijing"]
	if isPresent {
		fmt.Printf("The value of \"Beijing\" in map1 is: %d\n", value)
	} else {
		fmt.Printf("map1 does not contain Beijing")
	}
	value, isPresent = map1["sb"]
	fmt.Println(isPresent)
}