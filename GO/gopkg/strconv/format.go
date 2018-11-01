package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.FormatFloat(10.111, 'e', 5, 64))
	fmt.Println(strconv.FormatInt(12312324234343433, 2))
	fmt.Println(strconv.FormatUint(1231232423434343322, 2))
}
