package main

import (
	"fmt"
	"time"
)

func tick(ch <-chan time.Time) {
	for t := range ch {
		fmt.Println(t)
	}
}

func main() {
	ticker := time.NewTicker(time.Second)

	fmt.Println(ticker)
	go tick(ticker.C)

	time.Sleep(5 * time.Second)
}
