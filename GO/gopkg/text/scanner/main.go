package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	src := strings.NewReader("int hello = 3; hello+23; print hello;")

	fmt.Println(src)

	var s scanner.Scanner
	s.Init(src)

	tok := s.Scan()
	fmt.Println(s.TokenText())
	for tok != scanner.EOF {
		tok = s.Scan()
		fmt.Println(s.TokenText())
	}
}
