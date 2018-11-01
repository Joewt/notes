package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Sum(n ...int) int {
	var s int
	for _, i := range n {
		s += i
	}
	return s
}

func main() {
	h := md5.New()
	io.WriteString(h, "joe")
	fmt.Printf("%x\n", h.Sum(nil))
	fmt.Println(Sum(1, 2, 3))
}
