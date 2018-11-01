package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	reader1, _ := os.Open("1.txt")
	reader2, _ := os.Open("2.txt")
	//limitReader := io.LimitReader(reader, 20)
	limitReader := io.MultiReader(reader1, reader2)
	fmt.Println(reflect.TypeOf(limitReader))
	p := make([]byte, 2)
	var n, total int
	var err error
	for {
		n, err = limitReader.Read(p)
		if err == io.EOF {
			fmt.Println("total", n)
			break
		}

		total = total + n
		fmt.Println("read value: ", string(p[0:n]), " read count: ", n)
	}
}
