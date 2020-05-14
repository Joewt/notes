package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello, how is it going, Hugo?"
	var manyG = "ggggggggg"
	fmt.Printf("%d\n", strings.Count(str, "H"))
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))
}