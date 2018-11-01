package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	src, _ := os.Open("src.txt")
	des, _ := os.Create("des.txt")
	//written, err := io.Copy(des, src)

	written, err := io.CopyN(des, src, 10)

	if err == nil {
		fmt.Println("copy success! ", written)
	}
}
