package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("abcasdfas122223123$sdfsdf"))
	r := bufio.NewReader(rb)
	line, err := r.ReadString('$')
	if err == nil {
		fmt.Println(line)
	}

}
