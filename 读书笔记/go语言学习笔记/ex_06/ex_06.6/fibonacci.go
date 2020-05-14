package main

import "fmt"

func main() {
	result := 0
	for i := 0; i <= 100; i++ {
		result = fibonacci(i)
		fmt.Println(result)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1{
		res =1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}