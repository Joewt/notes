package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	n, err := strconv.Atoi("123223423423")
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(n))
}
