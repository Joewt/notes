package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 3)
	len := copy(dst, src[4:])
	fmt.Println(dst, len)
	dst = append(dst, 1)
	fmt.Println(dst)
}
