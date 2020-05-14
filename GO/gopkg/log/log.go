package main

import (
	"fmt"
	"log"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err: ", err)
			log.Println("recover...")
		}
	}()

	log.Panic("panic and stop")
	fmt.Println("not to called")
}
