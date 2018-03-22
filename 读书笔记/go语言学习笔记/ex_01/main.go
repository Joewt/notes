package main

import "fmt"

func main() {
	ch := make(chan int)

	go loop(ch)
	go loop(ch)

	<-ch
	<-ch

}

func loop(ch chan int) {
	for i := 0; i < 8; i++ {
		fmt.Println(i)
	}
	ch <- 1
}