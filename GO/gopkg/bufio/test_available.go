package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wb)
	fmt.Println(w.Available())
	fmt.Println(len(wb.Bytes()), string(wb.Bytes()))
	w.WriteByte('1')
	fmt.Println(w.Available())
	w.Flush()
	fmt.Println(w.Available())
	fmt.Println(len(wb.Bytes()), string(wb.Bytes()))
}
