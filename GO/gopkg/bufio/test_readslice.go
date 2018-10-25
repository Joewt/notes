package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("12345$67890"))
	r := bufio.NewReader(rb)
	line, err := r.ReadSlice('$')
	if err != nil {
		return
	}

	fmt.Println(string(line))
	r.ReadSlice('$')
	fmt.Println(string(line))

}
