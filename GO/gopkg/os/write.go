package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("hello.go", os.O_RDWR, 0420)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b := make([]byte, 10)
	fi, _ := f.Stat()
	fmt.Println("size : ", fi.Size())
	n, _ := f.Read(b)
	fmt.Printf("content: %s\n", b[:n])

	f.WriteString(" world\n")
	f.Seek(0, 0)
	n, _ = f.Read(b)
	fmt.Printf("content: %s\n", b[:n])

}
