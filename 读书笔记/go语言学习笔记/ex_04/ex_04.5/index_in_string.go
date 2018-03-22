package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hi, I'm Marc, Hi."
	var str1 = strings.Replace(str, "Hi", "sb", 2)
	fmt.Println(str1)
	fmt.Printf("The position of \"Marc\" is:")
	fmt.Printf("%d\n", strings.Index(str, "Marc"))
}