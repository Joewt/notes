package main

import (
	"fmt"
	"os"
)

func main() {
	if err := os.Symlink("joe.go", "joewt.go"); err != nil {
		panic(err)
	}
	fmt.Println("link created success!!")
}
