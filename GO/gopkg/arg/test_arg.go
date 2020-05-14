package main

import (
	"flag"
	"fmt"
)

var boolFlag bool

func init() {
	flag.BoolVar(&boolFlag, "flag", false, "help flag")
}

func main() {
	flag.Parse()
	fmt.Println("boolFlag: ", boolFlag)
}
