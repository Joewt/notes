package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("12345,67890$"))
	r := bufio.NewReader(rb)
	line, _ := r.ReadSlice(',')
	fmt.Println(string(line))
	r.UnreadByte()
	line, _ = r.ReadSlice('$')
	fmt.Println(string(line))
}
