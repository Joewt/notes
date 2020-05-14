package main

import (
	"fmt"
	"os"
)

var (
	c, d int
	err  error
)

func main() {
	fmt.Printf("%q\n", "加引号")
	fmt.Printf("%q\n", 65)
	var a, b int
	n, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		fmt.Println("Scan error")
	}

	fmt.Fprintln(os.Stdout, n, a, b)

	n, err = fmt.Fscan(os.Stdin, &c, &d)
	fmt.Printf("%v %v", n, err)

}
