package main

import (
	"fmt"
	"sort"
)

func main() {
	a := sort.IntSlice{5, 2, 6, 3, 1, 4} // unsorted

	fmt.Println(a.Len()) // 6

	fmt.Println(a.Less(0, 1)) // false

	fmt.Println(a.Search(4)) // 5

	a.Swap(0, 1)
	fmt.Println(a) // Output: [2 5 6 3 1 4]

	a.Sort()
	fmt.Println(a) // Output: [1 2 3 4 5 6]
	sort.Sort(sort.Reverse(a))
	fmt.Println(a)
}
