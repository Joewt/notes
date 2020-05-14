package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getegid())
	fmt.Println(os.Geteuid())
	fmt.Println(os.Getpagesize())
	fmt.Println(os.Getpid())

	fmt.Println(os.Getppid())
	fmt.Println(os.Getuid())
	pwd, _ := os.Getwd()

	fmt.Println(pwd)
	name, _ := os.Hostname()
	fmt.Println(name)
}
