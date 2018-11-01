package main

import (
	"fmt"
	"sync"
	"time"
)

func waiter(cond *sync.Cond, id int) {
	cond.L.Lock()
	cond.Wait()
	cond.L.Unlock()
	fmt.Printf("Waiter: %d wake up \n", id)
}

func main() {

	locker := new(sync.Mutex)
	cond := sync.NewCond(locker)
	for i := 0; i < 3; i++ {
		go waiter(cond, i)
	}

	time.Sleep(time.Second * 1)
	cond.L.Lock()
	cond.Signal()
	cond.L.Unlock()
	for i := 3; i < 5; i++ {
		go waiter(cond, i)
	}

	time.Sleep(time.Second * 1)

	cond.L.Lock()
	cond.Signal()
	cond.L.Unlock()

	cond.L.Lock()
	cond.Broadcast()
	cond.L.Unlock()
	time.Sleep(time.Second * 1)
}
