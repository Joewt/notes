package main

import (
	"fmt"
	"runtime"
	"sync"
)

func click(total *int, ch chan int) {
	for i := 0; i < 100000; i++ {
		*total += 1

	}
	ch <- 1
}

func clickmutex(total *int, ch chan int, m *sync.Mutex) {
	for i := 0; i < 100; i++ {
		m.Lock()
		*total += 1
		m.Unlock()
	}
	ch <- 1
}

func main() {
	runtime.GOMAXPROCS(2)

	m := new(sync.Mutex)

	count1 := 0
	count2 := 0
	ch := make(chan int, 200)

	for i := 0; i < 1000; i++ {
		go click(&count1, ch)
	}
	for i := 0; i < 200; i++ {
		go clickmutex(&count2, ch, m)
	}

	for i := 0; i < 200; i++ {
		<-ch
	}
	fmt.Printf("%d\n%d\n", count1, count2)
}
