package main

import (
	"fmt"
	"unicode/utf8"
)

func main() { 
	str1 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("%d\n", len(str1))
	fmt.Printf("%d\n", utf8.RuneCountInString(str1))
}