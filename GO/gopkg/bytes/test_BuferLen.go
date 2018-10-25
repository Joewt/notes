package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer([]byte("1234567.89"))
	fmt.Println(b.Len())
	fmt.Println(string(b.Next(3)))
	line, err := b.ReadBytes('.')
	if err == nil {
		fmt.Println(string(line), err)
	}
	fmt.Println(b.Len())
	fmt.Println(b.String())
	b.Reset()
	fmt.Println(b.Len())
}
