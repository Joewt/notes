package main

import (
	"bytes"
	"fmt"
)

func print(s [][]byte) {
	if len(s) == 0 {
		fmt.Println("nil")
		return
	}
	for i, c := range s {
		fmt.Println(i, string(c))
	}
}

func main() {
	b := bytes.NewBuffer(nil)
	fmt.Println(b.Write([]byte("12345")))
	fmt.Println(string(b.Bytes()))
	b.WriteByte('8')
	fmt.Println(string(b.Bytes()))
	b.WriteRune('二')
	fmt.Println(string(b.Bytes()))
	b.WriteString("鱼")
	fmt.Println(string(b.Bytes()))
	c := []byte("二鱼")
	c = bytes.Repeat(c, 3)
	d := bytes.Replace(c, []byte("二"), []byte("傻"), -1)
	fmt.Println(string(d))

	str := []byte("你好")
	fmt.Println(len(str))
	s := bytes.Runes(str)
	fmt.Println(len(s))
	for i, c := range bytes.Split(str, nil) {
		fmt.Println(i, string(c))
	}
	print(bytes.SplitAfter(str, nil))
	ss := []byte("hello world")
	fmt.Println(string(bytes.Title(ss)))
	fmt.Println(string(bytes.ToLower(ss)))
	fmt.Println(string(bytes.ToUpper(ss)))
	sss := []byte(" 123 4 124 ")
	fmt.Println(string(bytes.Trim(sss, " ")))

}
