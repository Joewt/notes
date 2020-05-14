package main

import "fmt"

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 2
}

func main() {
	fmt.Println(f())
}