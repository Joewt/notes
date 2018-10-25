package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("hello world"))
	wb := bytes.NewBuffer(nil)
	r := bufio.NewReader(rb)
	w := bufio.NewWriter(wb)
	rw := bufio.NewReadWriter(r, w)
	var rbuf [128]byte
	if n, err := rw.Read(rbuf[:]); err != nil {
		return
	} else {
		fmt.Println(string(rbuf[:n]))
	}

	if _, err := rw.Write([]byte("hello")); err != nil {
		return
	}
	rw.Flush()
	fmt.Println(string(wb.Bytes()))

}
