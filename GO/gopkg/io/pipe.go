package main

import (
	"fmt"
	"io"
)

func main() {
	r, w := io.Pipe()
	data := []byte("1jsldkjflasdkjfla;sdkf")
	go w.Write(data)
	outData := make([]byte, 10)
	n, _ := r.Read(outData)
	fmt.Println(string(outData), n)
}
