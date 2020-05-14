package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wb)
	w.Write([]byte("hello world"))
	fmt.Printf("%d: %s\n", len(wb.Bytes()), string(wb.Bytes()))
	w.Flush()
	fmt.Printf("%d: %s\n", len(wb.Bytes()), string(wb.Bytes()))

	rb := bytes.NewBuffer([]byte("hello"))
	r := bufio.NewReaderSize(rb, 10)
	var buf [10]byte
	n, _ := r.Read(buf[:])
	fmt.Println(string(buf[:n]))

}
