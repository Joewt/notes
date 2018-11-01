package main

import (
	"fmt"
	"strconv"
)

func main() {
	newList := strconv.AppendBool(make([]byte, 3), false)
	fmt.Println(newList)
}
