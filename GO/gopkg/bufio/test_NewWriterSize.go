package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriterSize(wb, 12)
	w.Write([]byte("hello"))
	w.Write([]byte("world"))
	w.Flush()
	fmt.Println(len(wb.Bytes()), string(wb.Bytes()))

}
