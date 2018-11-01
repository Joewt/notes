package main

import (
	"fmt"
	"net"
)

func main() {
	mask := net.CIDRMask(2, 32)
	fmt.Printf("%#V", mask)
}
