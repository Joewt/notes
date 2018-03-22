package main

import "fmt"

func main() {
	n1 := number1(7)
	fmt.Println(n1(8))
	n2 := number2()
	n2()
}

func number1(x int) func(int) int {
	return func(y int) int {
		return x + y;
	}
}
func number2() func() {
	i := "Hello World"
	return func() {
		fmt.Println(i)
	}
}