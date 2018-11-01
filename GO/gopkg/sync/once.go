package main

import (
	"fmt"
	"sync"
)

func main() {
	once := new(sync.Once)
	ch := make(chan int, 3)
	for i := 0; i < 3; i++ {
		go func(x int) {
			once.Do(func() {
				fmt.Printf("once %d\n", x)
			})
			fmt.Printf("%d\n", x)
			ch <- 1
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-ch
	}
}
