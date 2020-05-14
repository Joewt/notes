package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 3, 1, 6, 10, 3, 9}
	sort.Ints(arr)
	fmt.Println(arr)
}
