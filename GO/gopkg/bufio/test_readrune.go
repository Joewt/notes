package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("你好"))
	r := bufio.NewReader(rb)
	c, size, err := r.ReadRune()
	if err == nil {
		fmt.Println(string(c))
		fmt.Printf("%x %d\n", c, size)
	}
}
