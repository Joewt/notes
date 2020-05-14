package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("hello.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b := make([]byte, 15)
	n, _ := f.Read(b)
	fmt.Printf("read total: %d\ncontent: %s\n", n, b[:n])

	fi, err := os.Open("exec")
	if err != nil {
		panic(err)
	}

	defer fi.Close()
	ff, _ := fi.Readdir(10)
	for i, f := range ff {
		fmt.Printf("file %d: %+V\n", i, f)
	}

}
