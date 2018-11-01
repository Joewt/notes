package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}

	addr1 := net.JoinHostPort("joe", "8080")
	fmt.Println(addr1)
	fmt.Println(addr)
	addr2, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(addr2)
	host, port, err := net.SplitHostPort("www.hnitoj.cn:80")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(host, port)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	ip := net.ParseIP("192.168.1.1")
	fmt.Printf("%#v\n", ip)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	names, err := net.LookupIP("www.hnitoj.com")
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println(names)
	}

}
