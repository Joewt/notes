package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("chdir.go")
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v", err)
		return
	}
	defer f.Close()
	b := make([]byte, 2000)
	f.Seek(12, 0)
	n, _ := f.Read(b)
	fmt.Printf("%s\n", b[:n])
}
