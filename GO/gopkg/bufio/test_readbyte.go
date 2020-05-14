package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("12345$6"))
	//r := bufio.NewReader(rb)
	//b, err := r.ReadByte()
	//fmt.Println(string(b), err)

	r1 := bufio.NewReader(rb)

	b1, err := r1.ReadBytes('$')
	fmt.Println(string(b1), err)

}
