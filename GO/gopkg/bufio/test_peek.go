package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("hello1234567890"))
	r := bufio.NewReader(rb)
	b1, _ := r.Peek(5)
	fmt.Println(string(b1))

}
