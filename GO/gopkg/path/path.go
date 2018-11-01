package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Base("/a/b/"))
	fmt.Println(path.Base(""))
	fmt.Println(path.Base("/"))
}
