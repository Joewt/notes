package main

import (
	"container/heap"
	"fmt"
)

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h myHeap) printHeap() {
	n := 1
	level := 1
	for n <= h.Len() {
		fmt.Println(h[n-1 : n-1+level])
		n += level
		level *= 2
	}
}

func main() {
	data := []int{12, 43, 1, 4, 88, 90, 33, 55, 11, 12, 34, 111, 99, 110, 89}
	aheap := new(myHeap)
	for _, v := range data {
		aheap.Push(v)
	}
	aheap.printHeap()
	heap.Init(aheap)
	fmt.Println(*aheap)

}
