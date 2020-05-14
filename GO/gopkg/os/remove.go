package main

import (
	"fmt"
	"os"
)

func main() {
	//err := os.Remove("hello.go")
	err := os.RemoveAll("hello/hello")
	if err != nil {
		panic(err)
	}

	fmt.Println("remove success!")
}
