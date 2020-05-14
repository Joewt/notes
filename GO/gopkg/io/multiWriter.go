package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	w1, _ := os.Create("a.txt")
	w2, _ := os.Create("b.txt")
	mutiWriter := io.MultiWriter(w1, w2)
	p := []byte("sfsafasfasfsa")
	n, err := mutiWriter.Write(p)
	if err == nil {
		fmt.Println("write total ", n)
	}
}
