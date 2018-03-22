package main

import "fmt"

func main() {
	fmt.Printf("%d\n", Mu(2, 3, 4))
}
func Mu(a int, b int, c int) int {
	return a * b * c
}