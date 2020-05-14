package main

import (
	"flag"
	"fmt"
)

var f = flag.String("f", "abc", "test f")

func main() {
	flag.Parse()
	fmt.Println(*f)
	fmt.Println(flag.Args())
}
