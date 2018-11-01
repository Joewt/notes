package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	r, err := os.Open("src.txt")
	if err != nil {
		panic(err)
	}

	var n, total int

	p := make([]byte, 10)

	for {
		n, err = io.ReadAtLeast(r, p, 10)
		if err == nil {
			fmt.Println("read enough value")

		}
		if err == io.EOF {
			fmt.Println("read eof")
			break
		}
		total = total + n
		fmt.Println(string(p[0:n]))
	}
}
