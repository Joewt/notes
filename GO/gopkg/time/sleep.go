package main

import (
	"fmt"
	"time"
)

func main() {
	time.Sleep(1000 * time.Millisecond)
	time.Sleep(1 * time.Second)
	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Println(now)
	}
}
