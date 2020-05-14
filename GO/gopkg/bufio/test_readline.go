package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("123\r\nsdfasldfjs\r\nsdfsf\r\n"))
	r := bufio.NewReader(rb)
	line, prefix, err := r.ReadLine()
	if err == nil {
		fmt.Println(line, string(line), prefix)
	}
}
