package main

import (
	"fmt"
	"net/http"
)

func main() {
	var notingData []byte = []byte{0x01, 0x02}
	var gifData []byte = []byte{0x01, 0x02}
	fmt.Println(http.DetectContentType(notingData))
	fmt.Println(http.DetectContentType(gifData))
}
