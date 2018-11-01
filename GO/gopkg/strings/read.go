package main

import (
	"fmt"
	"strings"
)

func main() {
	read := strings.NewReader("a bcdef 呦呦呦")
	var b []byte
	fmt.Println(read.Len())
	b = make([]byte, 10)
	n, err := read.Read(b)
	if err != nil {
		panic(err)
	}

	fmt.Println(n, b)
}
