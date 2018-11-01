package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(t.Intn(100))
}
