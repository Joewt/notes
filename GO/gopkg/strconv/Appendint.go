package main

import (
	"fmt"
	"strconv"
)

func main() {
	newList := strconv.AppendInt(make([]byte, 0), 123000, 10)

	fmt.Println(newList)
	newlist := strconv.AppendQuoteRune(make([]byte, 0), 'a')
	fmt.Println(newlist)
	newlist = strconv.AppendQuoteRune(make([]byte, 0), '\'')
	fmt.Println(newlist)
	newlist = strconv.AppendQuoteRune(make([]byte, 0), '中')
	fmt.Println(newlist)
}
