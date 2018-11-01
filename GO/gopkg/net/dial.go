package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.hnitoj.cn:80")
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%#v", conn)
}
