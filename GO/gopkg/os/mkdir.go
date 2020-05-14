package main

import "os"

func main() {
	if err := os.Mkdir("hello_world", 0777); err != nil {
		panic(err)
	}

}
