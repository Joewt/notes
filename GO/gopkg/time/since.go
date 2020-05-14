package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	time.Sleep(3 * time.Second)
	fmt.Println(time.Since(t))
}
