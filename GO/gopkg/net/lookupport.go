package main

import (
	"fmt"
	"net"
)

func main() {
	port, err := net.LookupPort("tcp", "https")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(port)
}
